package just

import (
	"github.com/danztran/benchmark-protobuf-marshalers/pb"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"google.golang.org/protobuf/proto"
)

func protoClone[T proto.Message](d T) T {
	return proto.Clone(d).(T)
}

func genProduct() proto.Message {
	d := &pb.Product{}
	err := faker.FakeData(d,
		options.WithGenerateUniqueValues(true),
		options.WithRandomMapAndSliceMaxSize(10))
	if err != nil {
		panic(err)
	}
	return d
}

func newProduct() proto.Message {
	return new(pb.Product)
}

func genReview() proto.Message {
	d := &pb.Product_Review{}
	err := faker.FakeData(d,
		options.WithGenerateUniqueValues(true),
		options.WithRandomMapAndSliceMaxSize(10))
	if err != nil {
		panic(err)
	}
	return d
}

func newReview() proto.Message {
	return new(pb.Product_Review)
}

func genShipping() proto.Message {
	d := &pb.Product_Shipping_Address{}
	err := faker.FakeData(d,
		options.WithGenerateUniqueValues(true),
		options.WithRandomMapAndSliceMaxSize(10))
	if err != nil {
		panic(err)
	}
	return d
}

func newShipping() proto.Message {
	return new(pb.Product_Shipping_Address)
}
