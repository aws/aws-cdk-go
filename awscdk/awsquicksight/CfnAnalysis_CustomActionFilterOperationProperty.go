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
type CfnAnalysis_CustomActionFilterOperationProperty struct {
	// `CfnAnalysis.CustomActionFilterOperationProperty.SelectedFieldsConfiguration`.
	SelectedFieldsConfiguration interface{} `field:"required" json:"selectedFieldsConfiguration" yaml:"selectedFieldsConfiguration"`
	// `CfnAnalysis.CustomActionFilterOperationProperty.TargetVisualsConfiguration`.
	TargetVisualsConfiguration interface{} `field:"required" json:"targetVisualsConfiguration" yaml:"targetVisualsConfiguration"`
}

