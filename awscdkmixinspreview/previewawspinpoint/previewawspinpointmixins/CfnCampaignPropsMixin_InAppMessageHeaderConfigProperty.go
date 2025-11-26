package previewawspinpointmixins


// Specifies the configuration and content of the header or title text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inAppMessageHeaderConfigProperty := &InAppMessageHeaderConfigProperty{
//   	Alignment: jsii.String("alignment"),
//   	Header: jsii.String("header"),
//   	TextColor: jsii.String("textColor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessageheaderconfig.html
//
type CfnCampaignPropsMixin_InAppMessageHeaderConfigProperty struct {
	// The text alignment of the title of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessageheaderconfig.html#cfn-pinpoint-campaign-inappmessageheaderconfig-alignment
	//
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The header or title text of the in-app message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessageheaderconfig.html#cfn-pinpoint-campaign-inappmessageheaderconfig-header
	//
	Header *string `field:"optional" json:"header" yaml:"header"`
	// The color of the body text, expressed as a string consisting of a hex color code (such as "#000000" for black).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessageheaderconfig.html#cfn-pinpoint-campaign-inappmessageheaderconfig-textcolor
	//
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

