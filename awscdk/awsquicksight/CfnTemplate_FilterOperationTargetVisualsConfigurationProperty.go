package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterOperationTargetVisualsConfigurationProperty := &FilterOperationTargetVisualsConfigurationProperty{
//   	SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   		TargetVisualOptions: jsii.String("targetVisualOptions"),
//   		TargetVisuals: []*string{
//   			jsii.String("targetVisuals"),
//   		},
//   	},
//   }
//
type CfnTemplate_FilterOperationTargetVisualsConfigurationProperty struct {
	// `CfnTemplate.FilterOperationTargetVisualsConfigurationProperty.SameSheetTargetVisualConfiguration`.
	SameSheetTargetVisualConfiguration interface{} `field:"optional" json:"sameSheetTargetVisualConfiguration" yaml:"sameSheetTargetVisualConfiguration"`
}

