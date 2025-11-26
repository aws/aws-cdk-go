package previewawspinpointmixins


// Specifies the configuration of main body text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inAppMessageBodyConfigProperty := &InAppMessageBodyConfigProperty{
//   	Alignment: jsii.String("alignment"),
//   	Body: jsii.String("body"),
//   	TextColor: jsii.String("textColor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagebodyconfig.html
//
type CfnCampaignPropsMixin_InAppMessageBodyConfigProperty struct {
	// The text alignment of the main body text of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagebodyconfig.html#cfn-pinpoint-campaign-inappmessagebodyconfig-alignment
	//
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The main body text of the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagebodyconfig.html#cfn-pinpoint-campaign-inappmessagebodyconfig-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The color of the body text, expressed as a string consisting of a hex color code (such as "#000000" for black).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagebodyconfig.html#cfn-pinpoint-campaign-inappmessagebodyconfig-textcolor
	//
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

