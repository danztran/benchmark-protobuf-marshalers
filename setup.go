package just

import (
	"github.com/danztran/benchmark-protobuf-marshalers/pb"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"google.golang.org/protobuf/proto"
)

var benchData = genProduct()

func genProduct() *pb.Product {
	d := newProduct()
	err := faker.FakeData(d,
		options.WithRandomMapAndSliceMaxSize(10))
	if err != nil {
		panic(err)
	}
	return d
}

func newProduct() *pb.Product {
	return new(pb.Product)
}

func protoClone[T proto.Message](d T) T {
	return proto.Clone(d).(T)
}
