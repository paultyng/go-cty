package ptypes

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/golang/protobuf/proto"
)

func TestRoundTripType(t *testing.T) {
	tests := []struct {
		Type cty.Type
	}{
		{cty.DynamicPseudoType},

		{cty.Number},
		{cty.String},
		{cty.Bool},

		{cty.List(cty.DynamicPseudoType)},
		{cty.List(cty.String)},
		{cty.Set(cty.DynamicPseudoType)},
		{cty.Set(cty.String)},
		{cty.Map(cty.DynamicPseudoType)},
		{cty.Map(cty.String)},

		{cty.EmptyTuple},
		{cty.Tuple([]cty.Type{cty.String})},
		{cty.Tuple([]cty.Type{cty.Bool})},

		{cty.EmptyObject},
		{cty.Object(map[string]cty.Type{"a": cty.Bool})},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%#v", test.Type), func(t *testing.T) {
			expected := MarshalType(test.Type)

			b, err := proto.Marshal(expected)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("encoded in %d bytes as %x", len(b), b)

			var actual Type
			err = proto.Unmarshal(b, &actual)
			if err != nil {
				t.Fatal(err)
			}

			got, err := UnmarshalType(&actual)
			if err != nil {
				t.Fatal(err)
			}

			if !got.Equals(test.Type) {
				t.Errorf(
					"value did not round-trip\ninput:  %#v\nresult: %#v",
					test.Type, got,
				)
			}
		})
	}
}
