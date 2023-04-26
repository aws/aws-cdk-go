package awsquicksight


// The options that style a section.
//
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
	// The height of a section.
	//
	// Heights can only be defined for header and footer sections. The default height margin is 0.5 inches.
	Height *string `field:"optional" json:"height" yaml:"height"`
	// The spacing between section content and its top, bottom, left, and right edges.
	//
	// There is no padding by default.
	Padding interface{} `field:"optional" json:"padding" yaml:"padding"`
}

