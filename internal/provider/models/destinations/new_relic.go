package destinations

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	. "github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/mezmo/terraform-provider-mezmo/internal/client"
	. "github.com/mezmo/terraform-provider-mezmo/internal/provider/models/modelutils"
)

type NewRelicDestinationModel struct {
	Id           String `tfsdk:"id"`
	PipelineId   String `tfsdk:"pipeline_id"`
	Title        String `tfsdk:"title"`
	Description  String `tfsdk:"description"`
	Inputs       List   `tfsdk:"inputs"`
	GenerationId Int64  `tfsdk:"generation_id"`
	AckEnabled   Bool   `tfsdk:"ack_enabled" user_config:"true"`
	Api          String `tfsdk:"api" user_config:"true"`
	AccountId    String `tfsdk:"account_id" user_config:"true"`
	LicenseKey   String `tfsdk:"license_key" user_config:"true"`
}

func NewRelicDestinationResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Represents a NewRelic destination.",
		Attributes: ExtendBaseAttributes(map[string]schema.Attribute{
			"api": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("logs"),
				Description: "New Relic API endpoint type",
				Validators:  []validator.String{stringvalidator.OneOf("logs", "metrics")},
			},
			"account_id": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "New Relic Account ID",
				Validators:  []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"license_key": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "New Relic License Key",
				Validators:  []validator.String{stringvalidator.LengthAtLeast(1)},
			},
		}, nil),
	}
}

func NewRelicDestinationFromModel(plan *NewRelicDestinationModel, previousState *NewRelicDestinationModel) (*Destination, diag.Diagnostics) {
	dd := diag.Diagnostics{}

	component := Destination{
		BaseNode: BaseNode{
			Type:        "new-relic",
			Title:       plan.Title.ValueString(),
			Description: plan.Description.ValueString(),
			Inputs:      StringListValueToStringSlice(plan.Inputs),
			UserConfig: map[string]any{
				"ack_enabled": plan.AckEnabled.ValueBool(),
				"api":         plan.Api.ValueString(),
				"account_id":  plan.AccountId.ValueString(),
				"license_key": plan.LicenseKey.ValueString(),
			},
		},
	}

	if previousState != nil {
		component.Id = previousState.Id.ValueString()
		component.GenerationId = previousState.GenerationId.ValueInt64()
	}

	return &component, dd
}

func NewRelicDestinationToModel(plan *NewRelicDestinationModel, component *Destination) {
	plan.Id = StringValue(component.Id)
	if component.Title != "" {
		plan.Title = StringValue(component.Title)
	}
	if component.Description != "" {
		plan.Description = StringValue(component.Description)
	}
	plan.GenerationId = Int64Value(component.GenerationId)
	plan.Inputs = SliceToStringListValue(component.Inputs)
	plan.AckEnabled = BoolValue(component.UserConfig["ack_enabled"].(bool))
	plan.Api = StringValue(component.UserConfig["api"].(string))
	plan.AccountId = StringValue(component.UserConfig["account_id"].(string))
	plan.LicenseKey = StringValue(component.UserConfig["license_key"].(string))
}
