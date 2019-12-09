package ptypes

import (
	"math/big"
	"sort"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
)

// MarshalValue produces a protobuf message of a cty.Value instance.
//
// The given value must conform to the given type, or an error will
// be returned.
func MarshalValue(val cty.Value, ty cty.Type) (*Value, error) {
	errs := val.Type().TestConformance(ty)
	if errs != nil {
		// Attempt a conversion
		var err error
		val, err = convert.Convert(val, ty)
		if err != nil {
			return nil, err
		}
	}

	// From this point onward, val can be assumed to be conforming to t.
	var path cty.Path
	v, err := marshalValue(val, path)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func marshalValue(val cty.Value, path cty.Path) (*Value, error) {
	ty := val.Type()
	v := &Value{
		Type: MarshalType(ty),
	}
	switch {
	case !val.IsKnown():
		v.Kind = &Value_UnknownValue{
			UnknownValue: UnknownValue_UNKNOWN_VALUE,
		}
		return v, nil
	case val.IsNull():
		v.Kind = &Value_NullValue{
			NullValue: NullValue_NULL_VALUE,
		}
		return v, nil
	case ty.IsPrimitiveType():
		switch ty {
		case cty.Number:
			switch {
			case val.RawEquals(cty.PositiveInfinity):
				v.Kind = &Value_NumberValue{
					NumberValue: &NumberValue{
						Kind: &NumberValue_InfinityValue{
							InfinityValue: InfinityValue_POSITIVE_INFINITY,
						},
					},
				}
				return v, nil
			case val.RawEquals(cty.NegativeInfinity):
				v.Kind = &Value_NumberValue{
					NumberValue: &NumberValue{
						Kind: &NumberValue_InfinityValue{
							InfinityValue: InfinityValue_NEGATIVE_INFINITY,
						},
					},
				}
				return v, nil
			default:
				bf := val.AsBigFloat()
				if iv, acc := bf.Int64(); acc == big.Exact {
					v.Kind = &Value_NumberValue{
						NumberValue: &NumberValue{
							Kind: &NumberValue_IntValue{
								IntValue: iv,
							},
						},
					}
				} else if fv, acc := bf.Float64(); acc == big.Exact {
					v.Kind = &Value_NumberValue{
						NumberValue: &NumberValue{
							Kind: &NumberValue_DoubleValue{
								DoubleValue: fv,
							},
						},
					}
				} else {
					v.Kind = &Value_NumberValue{
						NumberValue: &NumberValue{
							Kind: &NumberValue_StringValue{
								StringValue: bf.Text('f', -1),
							},
						},
					}
				}

				return v, nil
			}
		case cty.String:
			v.Kind = &Value_StringValue{
				StringValue: val.AsString(),
			}
			return v, nil
		case cty.Bool:
			v.Kind = &Value_BoolValue{
				BoolValue: val.True(),
			}
			return v, nil
		}
	case ty.IsListType(), ty.IsSetType():
		values := make([]*Value, 0, val.LengthInt())
		it := val.ElementIterator()
		path := append(path, nil)
		for it.Next() {
			ek, ev := it.Element()
			path[len(path)-1] = cty.IndexStep{
				Key: ek,
			}
			evv, err := marshalValue(ev, path)
			if err != nil {
				return nil, err
			}
			values = append(values, evv)
		}

		switch {
		case ty.IsListType():
			v.Kind = &Value_ListValue{
				ListValue: &ListValue{
					Values: values,
				},
			}
			return v, nil
		case ty.IsSetType():
			v.Kind = &Value_SetValue{
				SetValue: &SetValue{
					Values: values,
				},
			}
			return v, nil
		}
	case ty.IsTupleType():
		values := make([]*Value, 0, val.LengthInt())
		it := val.ElementIterator()
		path := append(path, nil) // local override of 'path' with extra element
		i := 0
		for it.Next() {
			ek, ev := it.Element()
			path[len(path)-1] = cty.IndexStep{
				Key: ek,
			}
			evv, err := marshalValue(ev, path)
			if err != nil {
				return nil, err
			}
			values = append(values, evv)
			i++
		}

		v.Kind = &Value_TupleValue{
			TupleValue: &TupleValue{
				Values: values,
			},
		}
		return v, nil

	case ty.IsMapType():
		values := make(map[string]*Value, val.LengthInt())
		it := val.ElementIterator()
		path := append(path, nil) // local override of 'path' with extra element
		for it.Next() {
			ek, ev := it.Element()
			path[len(path)-1] = cty.IndexStep{
				Key: ek,
			}
			key := ek.AsString()
			evv, err := marshalValue(ev, path)
			if err != nil {
				return nil, err
			}
			values[key] = evv
		}
		v.Kind = &Value_MapValue{
			MapValue: &MapValue{
				Values: values,
			},
		}
		return v, nil
	case ty.IsObjectType():
		atys := ty.AttributeTypes()
		path := append(path, nil) // local override of 'path' with extra element

		names := make([]string, 0, len(atys))
		for k := range atys {
			names = append(names, k)
		}
		sort.Strings(names)

		attrs := make(map[string]*Value, len(atys))
		for _, k := range names {
			av := val.GetAttr(k)
			path[len(path)-1] = cty.GetAttrStep{
				Name: k,
			}
			pbv, err := marshalValue(av, path)
			if err != nil {
				return nil, err
			}
			attrs[k] = pbv
		}
		v.Kind = &Value_ObjectValue{
			ObjectValue: &ObjectValue{
				Values: attrs,
			},
		}
		return v, nil
	case ty.IsCapsuleType():
		return nil, path.NewErrorf("capsule types not supported for protobuf encoding")
	default:
		// should never happen
		return nil, path.NewErrorf("cannot protobuf-serialize %s", ty.FriendlyName())
	}
	return nil, path.NewErrorf("not implemented")
}
