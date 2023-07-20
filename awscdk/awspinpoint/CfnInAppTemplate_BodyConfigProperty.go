package awspinpoint


// Specifies the configuration of the main body text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodyConfigProperty := &BodyConfigProperty{
//   	Alignment: jsii.String("alignment"),
//   	Body: jsii.String("body"),
//   	TextColor: jsii.String("textColor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-bodyconfig.html
//
type CfnInAppTemplate_BodyConfigProperty struct {
	// The text alignment of the main body text of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-bodyconfig.html#cfn-pinpoint-inapptemplate-bodyconfig-alignment
	//
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The main body text of the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-bodyconfig.html#cfn-pinpoint-inapptemplate-bodyconfig-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The color of the body text, expressed as a hex color code (such as #000000 for black).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-bodyconfig.html#cfn-pinpoint-inapptemplate-bodyconfig-textcolor
	//
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

