package mixinsawspinpoint


// Specifies the appearance of an in-app message, including the message type, the title and body text, text and background colors, and the configurations of buttons that appear in the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var customConfig interface{}
//
//   campaignInAppMessageProperty := &CampaignInAppMessageProperty{
//   	Content: []interface{}{
//   		&InAppMessageContentProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BodyConfig: &InAppMessageBodyConfigProperty{
//   				Alignment: jsii.String("alignment"),
//   				Body: jsii.String("body"),
//   				TextColor: jsii.String("textColor"),
//   			},
//   			HeaderConfig: &InAppMessageHeaderConfigProperty{
//   				Alignment: jsii.String("alignment"),
//   				Header: jsii.String("header"),
//   				TextColor: jsii.String("textColor"),
//   			},
//   			ImageUrl: jsii.String("imageUrl"),
//   			PrimaryBtn: &InAppMessageButtonProperty{
//   				Android: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				DefaultConfig: &DefaultButtonConfigurationProperty{
//   					BackgroundColor: jsii.String("backgroundColor"),
//   					BorderRadius: jsii.Number(123),
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   					Text: jsii.String("text"),
//   					TextColor: jsii.String("textColor"),
//   				},
//   				Ios: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				Web: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   			},
//   			SecondaryBtn: &InAppMessageButtonProperty{
//   				Android: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				DefaultConfig: &DefaultButtonConfigurationProperty{
//   					BackgroundColor: jsii.String("backgroundColor"),
//   					BorderRadius: jsii.Number(123),
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   					Text: jsii.String("text"),
//   					TextColor: jsii.String("textColor"),
//   				},
//   				Ios: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				Web: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   			},
//   		},
//   	},
//   	CustomConfig: customConfig,
//   	Layout: jsii.String("layout"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigninappmessage.html
//
type CfnCampaignPropsMixin_CampaignInAppMessageProperty struct {
	// An array that contains configurtion information about the in-app message for the campaign, including title and body text, text colors, background colors, image URLs, and button configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigninappmessage.html#cfn-pinpoint-campaign-campaigninappmessage-content
	//
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// Custom data, in the form of key-value pairs, that is included in an in-app messaging payload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigninappmessage.html#cfn-pinpoint-campaign-campaigninappmessage-customconfig
	//
	CustomConfig interface{} `field:"optional" json:"customConfig" yaml:"customConfig"`
	// A string that describes how the in-app message will appear. You can specify one of the following:.
	//
	// - `BOTTOM_BANNER` – a message that appears as a banner at the bottom of the page.
	// - `TOP_BANNER` – a message that appears as a banner at the top of the page.
	// - `OVERLAYS` – a message that covers entire screen.
	// - `MOBILE_FEED` – a message that appears in a window in front of the page.
	// - `MIDDLE_BANNER` – a message that appears as a banner in the middle of the page.
	// - `CAROUSEL` – a scrollable layout of up to five unique messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigninappmessage.html#cfn-pinpoint-campaign-campaigninappmessage-layout
	//
	Layout *string `field:"optional" json:"layout" yaml:"layout"`
}

