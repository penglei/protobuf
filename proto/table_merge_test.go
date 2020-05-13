package proto_test

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
	tpb "github.com/golang/protobuf/proto/proto3_proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

func TestMergeWrappedValue(t *testing.T) {
	dst := tpb.Wrapped{Foo: &wrappers.BoolValue{Value: true}}
	src := tpb.Wrapped{Foo: &wrappers.BoolValue{Value: false}}
	proto.Merge(&dst, &src)
	if dst.Foo.Value {
		t.Errorf("merge wrapped bool value(false) failed")
	}
}

func TestMergeScalaValue(t *testing.T) {
	dst := tpb.Scalar{Foo: true}
	src := tpb.Scalar{Foo: false}
	proto.Merge(&dst, &src)
	if !dst.Foo {
		t.Errorf("scalar bool value(false) shouldn't be merged into")
	}
}
