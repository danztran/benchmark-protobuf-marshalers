package just

import (
	"encoding/json"
	"path"
	"slices"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestJson(t *testing.T) {
	t.Parallel()
	assert := func(t *testing.T, want proto.Message, got proto.Message) {
		equal := proto.Equal(want, got)
		// marshal to json and diff
		if !equal {
			expected, err := json.Marshal(want)
			require.NoError(t, err)
			actual, err := json.Marshal(got)
			require.NoError(t, err)
			require.JSONEq(t, string(expected), string(actual))
		}
	}

	for sample := range slices.Values(samples) {
		want := sample.Random()
		for engine := range slices.Values(engines) {
			t.Run(path.Join(sample.Name, engine.Name), func(t *testing.T) {
				t.Parallel()
				require.NotZero(t, want)

				raw, err := engine.Marshal(want)
				require.NoError(t, err)

				got := sample.New()
				err = engine.Unmarshal(raw, got)
				require.NoError(t, err)

				assert(t, want, got)
			})
		}

		t.Run(path.Join(sample.Name, "Copier"), func(t *testing.T) {
			t.Parallel()
			got := sample.New()
			err := copier.CopyWithOption(got, want, copier.Option{
				// DeepCopy: true,
			})
			require.NoError(t, err)
			assert(t, want, got)
		})

		t.Run(path.Join(sample.Name, "Copier"), func(t *testing.T) {
			t.Parallel()
			got := proto.Clone(want)
			assert(t, want, got)
		})
	}
}
