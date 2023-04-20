package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionFilterOperationProperty := &CustomActionFilterOperationProperty{
//   	SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
//   		SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   		SelectedFields: []*string{
//   			jsii.String("selectedFields"),
//   		},
//   	},
//   	TargetVisualsConfiguration: &FilterOperationTargetVisualsConfigurationProperty{
//   		SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   			TargetVisualOptions: jsii.String("targetVisualOptions"),
//   			TargetVisuals: []*string{
//   				jsii.String("targetVisuals"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_CustomActionFilterOperationProperty struct {
	// `CfnDashboard.CustomActionFilterOperationProperty.SelectedFieldsConfiguration`.
	SelectedFieldsConfiguration interface{} `field:"required" json:"selectedFieldsConfiguration" yaml:"selectedFieldsConfiguration"`
	// `CfnDashboard.CustomActionFilterOperationProperty.TargetVisualsConfiguration`.
	TargetVisualsConfiguration interface{} `field:"required" json:"targetVisualsConfiguration" yaml:"targetVisualsConfiguration"`
}

