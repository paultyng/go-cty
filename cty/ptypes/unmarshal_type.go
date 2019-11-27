package ptypes

import (
	"github.com/zclconf/go-cty/cty"
)

// UnmarshalType converts a protobuf message back to a cty.Type.
func UnmarshalType(ty *Type) (cty.Type, error) {
	var path cty.Path
	return unmarshalType(ty, path)
}

func unmarshalType(ty *Type, path cty.Path) (cty.Type, error) {
	switch tyk := ty.Kind.(type) {
	case *Type_DynamicType:
		return cty.DynamicPseudoType, nil
	case *Type_PrimitiveType:
		switch tyk.PrimitiveType {
		case PrimitiveType_STRING:
			return cty.String, nil
		case PrimitiveType_BOOL:
			return cty.Bool, nil
		case PrimitiveType_NUMBER:
			return cty.Number, nil
		}
	case *Type_CollectionType:
		eleTy, err := unmarshalType(tyk.CollectionType.Element, path)
		if err != nil {
			return cty.NilType, err
		}
		switch tyk.CollectionType.Kind {
		case CollectionKind_LIST:
			return cty.List(eleTy), nil
		case CollectionKind_MAP:
			return cty.Map(eleTy), nil
		case CollectionKind_SET:
			return cty.Set(eleTy), nil
		}
	case *Type_TupleType:
		eleTys := make([]cty.Type, 0, len(tyk.TupleType.Types))
		for _, pbty := range tyk.TupleType.Types {
			// TODO: update path
			eleTy, err := unmarshalType(pbty, path)
			if err != nil {
				return cty.NilType, err
			}
			eleTys = append(eleTys, eleTy)
		}
		return cty.Tuple(eleTys), nil
	case *Type_ObjectType:
		attrTypes := make(map[string]cty.Type, len(tyk.ObjectType.Attributes))
		for k, pbty := range tyk.ObjectType.Attributes {
			// TODO: update path
			attrTy, err := unmarshalType(pbty, path)
			if err != nil {
				return cty.NilType, err
			}
			attrTypes[k] = attrTy
		}
		return cty.Object(attrTypes), nil
	}
	return cty.NilType, path.NewErrorf("type of kind %T not supported", ty.Kind)
}
