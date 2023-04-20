package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionStyleProperty := &SectionStyleProperty{
//   	Height: jsii.String("height"),
//   	Padding: &SpacingProperty{
//   		Bottom: jsii.String("bottom"),
//   		Left: jsii.String("left"),
//   		Right: jsii.String("right"),
//   		Top: jsii.String("top"),
//   	},
//   }
//
type CfnDashboard_SectionStyleProperty struct {
	// `CfnDashboard.SectionStyleProperty.Height`.
	Height *string `field:"optional" json:"height" yaml:"height"`
	// `CfnDashboard.SectionStyleProperty.Padding`.
	Padding interface{} `field:"optional" json:"padding" yaml:"padding"`
}

