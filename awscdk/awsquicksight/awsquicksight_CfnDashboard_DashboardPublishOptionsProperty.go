package awsquicksight


// Dashboard publish options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardPublishOptionsProperty := &dashboardPublishOptionsProperty{
//   	adHocFilteringOption: &adHocFilteringOptionProperty{
//   		availabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	exportToCsvOption: &exportToCSVOptionProperty{
//   		availabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	sheetControlsOption: &sheetControlsOptionProperty{
//   		visibilityState: jsii.String("visibilityState"),
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

