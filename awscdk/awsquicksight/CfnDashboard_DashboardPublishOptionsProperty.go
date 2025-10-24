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
//   	DataQaEnabledOption: &DataQAEnabledOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	DataStoriesSharingOption: &DataStoriesSharingOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	ExecutiveSummaryOption: &ExecutiveSummaryOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	ExportToCsvOption: &ExportToCSVOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	ExportWithHiddenFieldsOption: &ExportWithHiddenFieldsOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	QuickSuiteActionsOption: &QuickSuiteActionsOptionProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html
//
type CfnDashboard_DashboardPublishOptionsProperty struct {
	// Ad hoc (one-time) filtering option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-adhocfilteringoption
	//
	AdHocFilteringOption interface{} `field:"optional" json:"adHocFilteringOption" yaml:"adHocFilteringOption"`
	// The drill-down options of data points in a dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-datapointdrillupdownoption
	//
	DataPointDrillUpDownOption interface{} `field:"optional" json:"dataPointDrillUpDownOption" yaml:"dataPointDrillUpDownOption"`
	// The data point menu label options of a dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-datapointmenulabeloption
	//
	DataPointMenuLabelOption interface{} `field:"optional" json:"dataPointMenuLabelOption" yaml:"dataPointMenuLabelOption"`
	// The data point tool tip options of a dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-datapointtooltipoption
	//
	DataPointTooltipOption interface{} `field:"optional" json:"dataPointTooltipOption" yaml:"dataPointTooltipOption"`
	// Adds Q&A capabilities to an Quick Sight dashboard.
	//
	// If no topic is linked, Dashboard Q&A uses the data values that are rendered on the dashboard. End users can use Dashboard Q&A to ask for different slices of the data that they see on the dashboard. If a topic is linked, Topic Q&A is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-dataqaenabledoption
	//
	DataQaEnabledOption interface{} `field:"optional" json:"dataQaEnabledOption" yaml:"dataQaEnabledOption"`
	// Data stories sharing option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-datastoriessharingoption
	//
	DataStoriesSharingOption interface{} `field:"optional" json:"dataStoriesSharingOption" yaml:"dataStoriesSharingOption"`
	// Executive summary option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-executivesummaryoption
	//
	ExecutiveSummaryOption interface{} `field:"optional" json:"executiveSummaryOption" yaml:"executiveSummaryOption"`
	// Export to .csv option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-exporttocsvoption
	//
	ExportToCsvOption interface{} `field:"optional" json:"exportToCsvOption" yaml:"exportToCsvOption"`
	// Determines if hidden fields are exported with a dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-exportwithhiddenfieldsoption
	//
	ExportWithHiddenFieldsOption interface{} `field:"optional" json:"exportWithHiddenFieldsOption" yaml:"exportWithHiddenFieldsOption"`
	// Determines if Actions in Amazon Quick Suite are enabled in a dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-quicksuiteactionsoption
	//
	QuickSuiteActionsOption interface{} `field:"optional" json:"quickSuiteActionsOption" yaml:"quickSuiteActionsOption"`
	// Sheet controls option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-sheetcontrolsoption
	//
	SheetControlsOption interface{} `field:"optional" json:"sheetControlsOption" yaml:"sheetControlsOption"`
	// The sheet layout maximization options of a dashbaord.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-sheetlayoutelementmaximizationoption
	//
	SheetLayoutElementMaximizationOption interface{} `field:"optional" json:"sheetLayoutElementMaximizationOption" yaml:"sheetLayoutElementMaximizationOption"`
	// The axis sort options of a dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-visualaxissortoption
	//
	VisualAxisSortOption interface{} `field:"optional" json:"visualAxisSortOption" yaml:"visualAxisSortOption"`
	// The menu options of a visual in a dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-visualmenuoption
	//
	VisualMenuOption interface{} `field:"optional" json:"visualMenuOption" yaml:"visualMenuOption"`
	// The visual publish options of a visual in a dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-visualpublishoptions
	//
	VisualPublishOptions interface{} `field:"optional" json:"visualPublishOptions" yaml:"visualPublishOptions"`
}

