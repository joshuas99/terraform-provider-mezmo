package destinations

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	. "github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	. "github.com/mezmo/terraform-provider-mezmo/internal/client"
)

type BlackholeDestinationModel struct {
	Id           String `tfsdk:"id"`
	PipelineId   String `tfsdk:"pipeline_id"`
	Title        String `tfsdk:"title"`
	Description  String `tfsdk:"description"`
	Inputs       List   `tfsdk:"inputs"`
	GenerationId Int64  `tfsdk:"generation_id"`
	AckEnabled   Bool   `tfsdk:"ack_enabled" user_config:"true"`
}

func BlackholeDestinationResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Represents a blackhole destination.",
		Attributes:  ExtendBaseAttributes(map[string]schema.Attribute{}, nil),
	}
}

func BlackholeDestinationFromModel(plan *BlackholeDestinationModel, previousState *BlackholeDestinationModel) (*Destination, diag.Diagnostics) {
	dd := diag.Diagnostics{}
	component := Destination{
		BaseNode: BaseNode{
			Type:        "blackhole",
			Title:       plan.Title.ValueString(),
			Description: plan.Description.ValueString(),
			UserConfig: map[string]any{
				"ack_enabled": plan.AckEnabled.ValueBool(),
			},
		},
	}

	if !plan.Inputs.IsUnknown() {
		inputs := make([]string, 0)
		for _, v := range plan.Inputs.Elements() {
			value, _ := v.(basetypes.StringValue)
			inputs = append(inputs, value.ValueString())
		}
		component.Inputs = inputs
	}

	if previousState != nil {
		component.Id = previousState.Id.ValueString()
		component.GenerationId = previousState.GenerationId.ValueInt64()
	}

	return &component, dd
}

func BlackholeDestinationToModel(plan *BlackholeDestinationModel, component *Destination) {
	plan.Id = StringValue(component.Id)
	if component.Title != "" {
		plan.Title = StringValue(component.Title)
	}
	if component.Description != "" {
		plan.Description = StringValue(component.Description)
	}
	plan.GenerationId = Int64Value(component.GenerationId)
	if component.Inputs != nil {
		inputs := make([]attr.Value, 0)
		for _, v := range component.Inputs {
			inputs = append(inputs, StringValue(v))
		}
		plan.Inputs = ListValueMust(StringType, inputs)
	}
	if component.UserConfig["ack_enabled"] != nil {
		value, _ := component.UserConfig["ack_enabled"].(bool)
		plan.AckEnabled = BoolValue(value)
	}
}
