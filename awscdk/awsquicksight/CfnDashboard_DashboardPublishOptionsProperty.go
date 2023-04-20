package awsquicksight


// Dashboard publish options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardPublishOptionsProperty := &DashboardPublishOptionsProperty{
//   	AdHocFilteringOption: &AdHocFilteringOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	DataPointDrillUpDownOption: &DataPointDrillUpDownOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	DataPointMenuLabelOption: &DataPointMenuLabelOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	DataPointTooltipOption: &DataPointTooltipOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	ExportToCsvOption: &ExportToCSVOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	ExportWithHiddenFieldsOption: &ExportWithHiddenFieldsOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	SheetControlsOption: &SheetControlsOptionProperty{
//   		VisibilityState: jsii.String("visibilityState"),
//   	},
//   	SheetLayoutElementMaximizationOption: &SheetLayoutElementMaximizationOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	VisualAxisSortOption: &VisualAxisSortOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	VisualMenuOption: &VisualMenuOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	VisualPublishOptions: &DashboardVisualPublishOptionsProperty{
//   		ExportHiddenFieldsOption: &ExportHiddenFieldsOptionProperty{
//   			AvailabilityStatus: jsii.String("availabilityStatus"),
//   		},
//   	},
//   }
//
type CfnDashboard_DashboardPublishOptionsProperty struct {
	// Ad hoc (one-time) filtering option.
	AdHocFilteringOption interface{} `field:"optional" json:"adHocFilteringOption" yaml:"adHocFilteringOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.DataPointDrillUpDownOption`.
	DataPointDrillUpDownOption interface{} `field:"optional" json:"dataPointDrillUpDownOption" yaml:"dataPointDrillUpDownOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.DataPointMenuLabelOption`.
	DataPointMenuLabelOption interface{} `field:"optional" json:"dataPointMenuLabelOption" yaml:"dataPointMenuLabelOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.DataPointTooltipOption`.
	DataPointTooltipOption interface{} `field:"optional" json:"dataPointTooltipOption" yaml:"dataPointTooltipOption"`
	// Export to .csv option.
	ExportToCsvOption interface{} `field:"optional" json:"exportToCsvOption" yaml:"exportToCsvOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.ExportWithHiddenFieldsOption`.
	ExportWithHiddenFieldsOption interface{} `field:"optional" json:"exportWithHiddenFieldsOption" yaml:"exportWithHiddenFieldsOption"`
	// Sheet controls option.
	SheetControlsOption interface{} `field:"optional" json:"sheetControlsOption" yaml:"sheetControlsOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.SheetLayoutElementMaximizationOption`.
	SheetLayoutElementMaximizationOption interface{} `field:"optional" json:"sheetLayoutElementMaximizationOption" yaml:"sheetLayoutElementMaximizationOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.VisualAxisSortOption`.
	VisualAxisSortOption interface{} `field:"optional" json:"visualAxisSortOption" yaml:"visualAxisSortOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.VisualMenuOption`.
	VisualMenuOption interface{} `field:"optional" json:"visualMenuOption" yaml:"visualMenuOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.VisualPublishOptions`.
	VisualPublishOptions interface{} `field:"optional" json:"visualPublishOptions" yaml:"visualPublishOptions"`
}

