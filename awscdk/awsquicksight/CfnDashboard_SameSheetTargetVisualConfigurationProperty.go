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
type CfnDashboard_SameSheetTargetVisualConfigurationProperty struct {
	// `CfnDashboard.SameSheetTargetVisualConfigurationProperty.TargetVisualOptions`.
	TargetVisualOptions *string `field:"optional" json:"targetVisualOptions" yaml:"targetVisualOptions"`
	// `CfnDashboard.SameSheetTargetVisualConfigurationProperty.TargetVisuals`.
	TargetVisuals *[]*string `field:"optional" json:"targetVisuals" yaml:"targetVisuals"`
}

