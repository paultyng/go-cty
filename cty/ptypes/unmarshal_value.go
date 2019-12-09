package ptypes

import (
	"github.com/zclconf/go-cty/cty"
)

// UnmarshalValue converts a protobuf message of a Value to a cty.Value
func UnmarshalValue(v *Value, ty cty.Type) (cty.Value, error) {
	var path cty.Path
	return unmarshalValue(v, ty, path)
}

func unmarshalValue(val *Value, ty cty.Type, path cty.Path) (cty.Value, error) {
	switch vk := val.Kind.(type) {
	case *Value_UnknownValue:
		impTy, err := unmarshalType(val.Type, path)
		if err != nil {
			return cty.NilVal, err
		}
		return cty.UnknownVal(impTy), nil
	case *Value_NullValue:
		impTy, err := unmarshalType(val.Type, path)
		if err != nil {
			return cty.NilVal, err
		}
		return cty.NullVal(impTy), nil
	case *Value_StringValue:
		return cty.StringVal(vk.StringValue), nil
	case *Value_NumberValue:
		switch nk := vk.NumberValue.Kind.(type) {
		case *NumberValue_InfinityValue:
			switch nk.InfinityValue {
			case InfinityValue_POSITIVE_INFINITY:
				return cty.PositiveInfinity, nil
			case InfinityValue_NEGATIVE_INFINITY:
				return cty.NegativeInfinity, nil
			}
		case *NumberValue_IntValue:
			return cty.NumberIntVal(nk.IntValue), nil
		case *NumberValue_DoubleValue:
			return cty.NumberFloatVal(nk.DoubleValue), nil
		case *NumberValue_StringValue:
			// TODO: wrap error?
			return cty.ParseNumberVal(nk.StringValue)
		}
	case *Value_BoolValue:
		return cty.BoolVal(vk.BoolValue), nil
	case *Value_ListValue:
		if len(vk.ListValue.Values) == 0 {
			return cty.ListValEmpty(ty.ElementType()), nil
		}
		types := make([]cty.Type, len(vk.ListValue.Values))
		for i := 0; i < len(vk.ListValue.Values); i++ {
			types[i] = ty.ElementType()
		}
		vals, err := unmarshalValues(vk.ListValue.Values, types, path)
		if err != nil {
			return cty.NilVal, err
		}
		return cty.ListVal(vals), nil
	case *Value_SetValue:
		if len(vk.SetValue.Values) == 0 {
			return cty.SetValEmpty(ty.ElementType()), nil
		}
		types := make([]cty.Type, len(vk.SetValue.Values))
		for i := 0; i < len(vk.SetValue.Values); i++ {
			types[i] = ty.ElementType()
		}
		vals, err := unmarshalValues(vk.SetValue.Values, types, path)
		if err != nil {
			return cty.NilVal, err
		}
		return cty.SetVal(vals), nil
	case *Value_MapValue:
		if len(vk.MapValue.Values) == 0 {
			return cty.MapValEmpty(ty.ElementType()), nil
		}
		vals := make(map[string]cty.Value, len(vk.MapValue.Values))
		path = append(path, nil)
		for k, v := range vk.MapValue.Values {
			path[len(path)-1] = cty.IndexStep{
				Key: cty.StringVal(k),
			}
			vv, err := unmarshalValue(v, ty.ElementType(), path)
			if err != nil {
				return cty.NilVal, path.NewError(err)
			}
			vals[k] = vv
		}
		return cty.MapVal(vals), nil
	case *Value_TupleValue:
		if len(vk.TupleValue.Values) == 0 {
			return cty.TupleVal(nil), nil
		}
		etys := ty.TupleElementTypes()
		if len(etys) != len(vk.TupleValue.Values) {
			return cty.NilVal, path.NewErrorf("a tuple of length %d is required", len(etys))
		}
		vals, err := unmarshalValues(vk.TupleValue.Values, etys, path)
		if err != nil {
			return cty.NilVal, err
		}
		return cty.TupleVal(vals), nil
	case *Value_ObjectValue:
		if len(vk.ObjectValue.Values) == 0 {
			return cty.ObjectVal(nil), nil
		}
		atys := ty.AttributeTypes()
		if len(atys) != len(vk.ObjectValue.Values) {
			return cty.NilVal, path.NewErrorf("an object with %d attributes is required (%d given)", len(atys), len(vk.ObjectValue.Values))
		}

		vals := make(map[string]cty.Value, len(vk.ObjectValue.Values))
		path = append(path, nil)
		for k, pbv := range vk.ObjectValue.Values {
			path[len(path)-1] = cty.IndexStep{
				Key: cty.StringVal(k),
			}
			aty, exists := atys[k]
			if !exists {
				return cty.NilVal, path.NewErrorf("unsupported attribute")
			}
			cv, err := unmarshalValue(pbv, aty, path)
			if err != nil {
				return cty.NilVal, err
			}
			vals[k] = cv
		}

		return cty.ObjectVal(vals), nil
	}

	return cty.NilVal, path.NewErrorf("not implemented")
}

func unmarshalValues(values []*Value, types []cty.Type, path cty.Path) ([]cty.Value, error) {
	vals := make([]cty.Value, 0, len(values))
	path = append(path, nil)
	for i, lv := range values {
		path[len(path)-1] = cty.IndexStep{
			Key: cty.NumberIntVal(int64(i)),
		}
		cv, err := unmarshalValue(lv, types[i], path)
		if err != nil {
			return nil, err
		}
		vals = append(vals, cv)
	}
	return vals, nil
}
