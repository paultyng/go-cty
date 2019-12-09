package ptypes

import (
	"github.com/zclconf/go-cty/cty"
)

// MarshalType produces a protobuf message with the cty.Type information.
func MarshalType(ty cty.Type) *Type {
	switch ty {
	case cty.DynamicPseudoType:
		return &Type{
			Kind: &Type_DynamicType{
				DynamicType: DynamicType_DYNAMIC_TYPE,
			},
		}
	case cty.Number:
		return &Type{
			Kind: &Type_PrimitiveType{
				PrimitiveType: PrimitiveType_NUMBER,
			},
		}
	case cty.Bool:
		return &Type{
			Kind: &Type_PrimitiveType{
				PrimitiveType: PrimitiveType_BOOL,
			},
		}
	case cty.String:
		return &Type{
			Kind: &Type_PrimitiveType{
				PrimitiveType: PrimitiveType_STRING,
			},
		}
	default:
		switch {
		case ty.IsListType():
			return &Type{
				Kind: &Type_CollectionType{
					CollectionType: &CollectionType{
						Kind:    CollectionKind_LIST,
						Element: MarshalType(ty.ElementType()),
					},
				},
			}
		case ty.IsSetType():
			return &Type{
				Kind: &Type_CollectionType{
					CollectionType: &CollectionType{
						Kind:    CollectionKind_SET,
						Element: MarshalType(ty.ElementType()),
					},
				},
			}
		case ty.IsMapType():
			return &Type{
				Kind: &Type_CollectionType{
					CollectionType: &CollectionType{
						Kind:    CollectionKind_MAP,
						Element: MarshalType(ty.ElementType()),
					},
				},
			}
		case ty.IsObjectType():
			atts := ty.AttributeTypes()
			types := make(map[string]*Type, len(atts))
			for k, at := range atts {
				types[k] = MarshalType(at)
			}
			return &Type{
				Kind: &Type_ObjectType{
					ObjectType: &ObjectType{
						Attributes: types,
					},
				},
			}
		case ty.IsTupleType():
			etypes := ty.TupleElementTypes()
			types := make([]*Type, 0, len(etypes))
			for _, et := range etypes {
				types = append(types, MarshalType(et))
			}
			return &Type{
				Kind: &Type_TupleType{
					TupleType: &TupleType{
						Types: types,
					},
				},
			}
		}
	}
	panic("not implemented")
}
