package previewawspinpointmixins


// Specifies the configuration and contents of an in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inAppMessageContentProperty := &InAppMessageContentProperty{
//   	BackgroundColor: jsii.String("backgroundColor"),
//   	BodyConfig: &InAppMessageBodyConfigProperty{
//   		Alignment: jsii.String("alignment"),
//   		Body: jsii.String("body"),
//   		TextColor: jsii.String("textColor"),
//   	},
//   	HeaderConfig: &InAppMessageHeaderConfigProperty{
//   		Alignment: jsii.String("alignment"),
//   		Header: jsii.String("header"),
//   		TextColor: jsii.String("textColor"),
//   	},
//   	ImageUrl: jsii.String("imageUrl"),
//   	PrimaryBtn: &InAppMessageButtonProperty{
//   		Android: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   		DefaultConfig: &DefaultButtonConfigurationProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BorderRadius: jsii.Number(123),
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   			Text: jsii.String("text"),
//   			TextColor: jsii.String("textColor"),
//   		},
//   		Ios: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   		Web: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   	},
//   	SecondaryBtn: &InAppMessageButtonProperty{
//   		Android: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   		DefaultConfig: &DefaultButtonConfigurationProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BorderRadius: jsii.Number(123),
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   			Text: jsii.String("text"),
//   			TextColor: jsii.String("textColor"),
//   		},
//   		Ios: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   		Web: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html
//
type CfnCampaignPropsMixin_InAppMessageContentProperty struct {
	// The background color for an in-app message banner, expressed as a hex color code (such as #000000 for black).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-backgroundcolor
	//
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Specifies the configuration of main body text in an in-app message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-bodyconfig
	//
	BodyConfig interface{} `field:"optional" json:"bodyConfig" yaml:"bodyConfig"`
	// Specifies the configuration and content of the header or title text of the in-app message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-headerconfig
	//
	HeaderConfig interface{} `field:"optional" json:"headerConfig" yaml:"headerConfig"`
	// The URL of the image that appears on an in-app message banner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-imageurl
	//
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// An object that contains configuration information about the primary button in an in-app message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-primarybtn
	//
	PrimaryBtn interface{} `field:"optional" json:"primaryBtn" yaml:"primaryBtn"`
	// An object that contains configuration information about the secondary button in an in-app message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagecontent.html#cfn-pinpoint-campaign-inappmessagecontent-secondarybtn
	//
	SecondaryBtn interface{} `field:"optional" json:"secondaryBtn" yaml:"secondaryBtn"`
}

