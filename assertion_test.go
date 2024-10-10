package just

import (
	"path"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestJson(t *testing.T) {
	t.Parallel()
	for engine := range slices.Values(engines) {
		for sample := range slices.Values(samples) {
			t.Run(path.Join(sample.Name, engine.Name), func(t *testing.T) {
				t.Parallel()
				data := sample.Random()
				require.NotZero(t, data)

				raw, err := engine.Marshal(data)
				require.NoError(t, err)

				d := sample.New()
				err = engine.Unmarshal(raw, d)
				require.NoError(t, err)

				equal := proto.Equal(data, d)
				require.True(t, equal, d)
			})
		}
	}
}
