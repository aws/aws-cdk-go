package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spacingProperty := &SpacingProperty{
//   	Bottom: jsii.String("bottom"),
//   	Left: jsii.String("left"),
//   	Right: jsii.String("right"),
//   	Top: jsii.String("top"),
//   }
//
type CfnDashboard_SpacingProperty struct {
	// `CfnDashboard.SpacingProperty.Bottom`.
	Bottom *string `field:"optional" json:"bottom" yaml:"bottom"`
	// `CfnDashboard.SpacingProperty.Left`.
	Left *string `field:"optional" json:"left" yaml:"left"`
	// `CfnDashboard.SpacingProperty.Right`.
	Right *string `field:"optional" json:"right" yaml:"right"`
	// `CfnDashboard.SpacingProperty.Top`.
	Top *string `field:"optional" json:"top" yaml:"top"`
}

