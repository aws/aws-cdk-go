package awsquicksight


// The configuration of target visuals that you want to be filtered.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
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
	// The configuration of the same-sheet target visuals that you want to be filtered.
	SameSheetTargetVisualConfiguration interface{} `field:"optional" json:"sameSheetTargetVisualConfiguration" yaml:"sameSheetTargetVisualConfiguration"`
}

