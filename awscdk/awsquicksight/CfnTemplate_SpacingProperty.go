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
type CfnTemplate_SpacingProperty struct {
	// `CfnTemplate.SpacingProperty.Bottom`.
	Bottom *string `field:"optional" json:"bottom" yaml:"bottom"`
	// `CfnTemplate.SpacingProperty.Left`.
	Left *string `field:"optional" json:"left" yaml:"left"`
	// `CfnTemplate.SpacingProperty.Right`.
	Right *string `field:"optional" json:"right" yaml:"right"`
	// `CfnTemplate.SpacingProperty.Top`.
	Top *string `field:"optional" json:"top" yaml:"top"`
}

