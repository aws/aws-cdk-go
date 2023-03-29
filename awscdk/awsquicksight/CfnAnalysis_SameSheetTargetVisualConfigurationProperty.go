package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sameSheetTargetVisualConfigurationProperty := &SameSheetTargetVisualConfigurationProperty{
//   	TargetVisualOptions: jsii.String("targetVisualOptions"),
//   	TargetVisuals: []*string{
//   		jsii.String("targetVisuals"),
//   	},
//   }
//
type CfnAnalysis_SameSheetTargetVisualConfigurationProperty struct {
	// `CfnAnalysis.SameSheetTargetVisualConfigurationProperty.TargetVisualOptions`.
	TargetVisualOptions *string `field:"optional" json:"targetVisualOptions" yaml:"targetVisualOptions"`
	// `CfnAnalysis.SameSheetTargetVisualConfigurationProperty.TargetVisuals`.
	TargetVisuals *[]*string `field:"optional" json:"targetVisuals" yaml:"targetVisuals"`
}

