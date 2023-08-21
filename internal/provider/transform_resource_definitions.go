package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	. "github.com/mezmo-inc/terraform-provider-mezmo/internal/provider/models/transforms"
)

func NewStringifyTransformResource() resource.Resource {
	return &TransformResource[StringifyTransformModel]{
		typeName:          "stringify",
		fromModelFunc:     StringifyTransformFromModel,
		toModelFunc:       StringifyTransformToModel,
		getIdFunc:         func(m *StringifyTransformModel) basetypes.StringValue { return m.Id },
		getPipelineIdFunc: func(m *StringifyTransformModel) basetypes.StringValue { return m.PipelineId },
		getSchemaFunc:     StringifyTransformResourceSchema,
	}
}
func NewCompactFieldsTransformResource() resource.Resource {
	return &TransformResource[CompactFieldsTransformModel]{
		typeName:          "compact_fields",
		fromModelFunc:     CompactFieldsTransformFromModel,
		toModelFunc:       CompactFieldsTransformToModel,
		getIdFunc:         func(m *CompactFieldsTransformModel) basetypes.StringValue { return m.Id },
		getPipelineIdFunc: func(m *CompactFieldsTransformModel) basetypes.StringValue { return m.PipelineId },
		getSchemaFunc:     CompactFieldsTransformResourceSchema,
	}
}
