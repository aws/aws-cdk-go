package mixinsawspinpoint


// Specifies the configuration and content of the header or title text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   headerConfigProperty := &HeaderConfigProperty{
//   	Alignment: jsii.String("alignment"),
//   	Header: jsii.String("header"),
//   	TextColor: jsii.String("textColor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-headerconfig.html
//
type CfnInAppTemplatePropsMixin_HeaderConfigProperty struct {
	// The text alignment of the title of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-headerconfig.html#cfn-pinpoint-inapptemplate-headerconfig-alignment
	//
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The title text of the in-app message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-headerconfig.html#cfn-pinpoint-inapptemplate-headerconfig-header
	//
	Header *string `field:"optional" json:"header" yaml:"header"`
	// The color of the title text, expressed as a hex color code (such as #000000 for black).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-headerconfig.html#cfn-pinpoint-inapptemplate-headerconfig-textcolor
	//
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

