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
//   	ExportToCsvOption: &ExportToCSVOptionProperty{
//   		AvailabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	SheetControlsOption: &SheetControlsOptionProperty{
//   		VisibilityState: jsii.String("visibilityState"),
//   	},
//   }
//
type CfnDashboard_DashboardPublishOptionsProperty struct {
	// Ad hoc (one-time) filtering option.
	AdHocFilteringOption interface{} `field:"optional" json:"adHocFilteringOption" yaml:"adHocFilteringOption"`
	// Export to .csv option.
	ExportToCsvOption interface{} `field:"optional" json:"exportToCsvOption" yaml:"exportToCsvOption"`
	// Sheet controls option.
	SheetControlsOption interface{} `field:"optional" json:"sheetControlsOption" yaml:"sheetControlsOption"`
}

