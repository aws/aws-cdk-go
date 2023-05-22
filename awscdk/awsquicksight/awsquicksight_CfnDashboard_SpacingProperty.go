package awsquicksight


// The configuration of spacing (often a margin or padding).
//
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
	// Define the bottom spacing.
	Bottom *string `field:"optional" json:"bottom" yaml:"bottom"`
	// Define the left spacing.
	Left *string `field:"optional" json:"left" yaml:"left"`
	// Define the right spacing.
	Right *string `field:"optional" json:"right" yaml:"right"`
	// Define the top spacing.
	Top *string `field:"optional" json:"top" yaml:"top"`
}

