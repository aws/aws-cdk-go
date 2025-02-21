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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sectionstyle.html
//
type CfnTemplate_SectionStyleProperty struct {
	// The height of a section.
	//
	// Heights can only be defined for header and footer sections. The default height margin is 0.5 inches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sectionstyle.html#cfn-quicksight-template-sectionstyle-height
	//
	Height *string `field:"optional" json:"height" yaml:"height"`
	// The spacing between section content and its top, bottom, left, and right edges.
	//
	// There is no padding by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sectionstyle.html#cfn-quicksight-template-sectionstyle-padding
	//
	Padding interface{} `field:"optional" json:"padding" yaml:"padding"`
}

