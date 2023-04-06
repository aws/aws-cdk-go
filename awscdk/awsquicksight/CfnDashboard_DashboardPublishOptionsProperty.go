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
	// The drill-down options of data points in a dashboard.
	DataPointDrillUpDownOption interface{} `field:"optional" json:"dataPointDrillUpDownOption" yaml:"dataPointDrillUpDownOption"`
	// The data point menu label options of a dashboard.
	DataPointMenuLabelOption interface{} `field:"optional" json:"dataPointMenuLabelOption" yaml:"dataPointMenuLabelOption"`
	// The data point tool tip options of a dashboard.
	DataPointTooltipOption interface{} `field:"optional" json:"dataPointTooltipOption" yaml:"dataPointTooltipOption"`
	// Export to .csv option.
	ExportToCsvOption interface{} `field:"optional" json:"exportToCsvOption" yaml:"exportToCsvOption"`
	// Determines if hidden fields are exported with a dashboard.
	ExportWithHiddenFieldsOption interface{} `field:"optional" json:"exportWithHiddenFieldsOption" yaml:"exportWithHiddenFieldsOption"`
	// Sheet controls option.
	SheetControlsOption interface{} `field:"optional" json:"sheetControlsOption" yaml:"sheetControlsOption"`
	// The sheet layout maximization options of a dashbaord.
	SheetLayoutElementMaximizationOption interface{} `field:"optional" json:"sheetLayoutElementMaximizationOption" yaml:"sheetLayoutElementMaximizationOption"`
	// The axis sort options of a dashboard.
	VisualAxisSortOption interface{} `field:"optional" json:"visualAxisSortOption" yaml:"visualAxisSortOption"`
	// The menu options of a visual in a dashboard.
	VisualMenuOption interface{} `field:"optional" json:"visualMenuOption" yaml:"visualMenuOption"`
	// The visual publish options of a visual in a dashboard.
	VisualPublishOptions interface{} `field:"optional" json:"visualPublishOptions" yaml:"visualPublishOptions"`
}

