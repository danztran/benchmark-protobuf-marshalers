package just

import (
	"testing"

	"github.com/danztran/benchmark-protobuf-marshalers/pb"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"google.golang.org/protobuf/proto"
)

func genProduct(tb testing.TB) *pb.Product {
	tb.Helper()
	d := newProduct()
	err := faker.FakeData(d,
		options.WithRandomMapAndSliceMaxSize(100))
	checkErr(tb, err)
	return d
}

func newProduct() *pb.Product {
	return new(pb.Product)
}

func protoClone[T proto.Message](d T) T {
	return proto.Clone(d).(T)
}

func checkErr(b testing.TB, err error) {
	b.Helper()
	if err != nil {
		b.Fatal(err)
	}
}

func preBench(b *testing.B, raw []byte, err error) {
	b.Helper()
	checkErr(b, err)
	b.SetBytes(int64(len(raw)))
	b.ResetTimer()
}
