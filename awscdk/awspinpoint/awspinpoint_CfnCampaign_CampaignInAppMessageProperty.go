package awspinpoint


// Specifies the appearance of an in-app message, including the message type, the title and body text, text and background colors, and the configurations of buttons that appear in the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customConfig interface{}
//
//   campaignInAppMessageProperty := &campaignInAppMessageProperty{
//   	content: []interface{}{
//   		&inAppMessageContentProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			bodyConfig: &inAppMessageBodyConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				body: jsii.String("body"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			headerConfig: &inAppMessageHeaderConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				header: jsii.String("header"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			imageUrl: jsii.String("imageUrl"),
//   			primaryBtn: &inAppMessageButtonProperty{
//   				android: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				defaultConfig: &defaultButtonConfigurationProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					borderRadius: jsii.Number(123),
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   					text: jsii.String("text"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				ios: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				web: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   			},
//   			secondaryBtn: &inAppMessageButtonProperty{
//   				android: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				defaultConfig: &defaultButtonConfigurationProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					borderRadius: jsii.Number(123),
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   					text: jsii.String("text"),
//   					textColor: jsii.String("textColor"),
//   				},
//   				ios: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   				web: &overrideButtonConfigurationProperty{
//   					buttonAction: jsii.String("buttonAction"),
//   					link: jsii.String("link"),
//   				},
//   			},
//   		},
//   	},
//   	customConfig: customConfig,
//   	layout: jsii.String("layout"),
//   }
//
type CfnCampaign_CampaignInAppMessageProperty struct {
	// An array that contains configurtion information about the in-app message for the campaign, including title and body text, text colors, background colors, image URLs, and button configurations.
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// Custom data, in the form of key-value pairs, that is included in an in-app messaging payload.
	CustomConfig interface{} `field:"optional" json:"customConfig" yaml:"customConfig"`
	// A string that describes how the in-app message will appear. You can specify one of the following:.
	//
	// - `BOTTOM_BANNER` – a message that appears as a banner at the bottom of the page.
	// - `TOP_BANNER` – a message that appears as a banner at the top of the page.
	// - `OVERLAYS` – a message that covers entire screen.
	// - `MOBILE_FEED` – a message that appears in a window in front of the page.
	// - `MIDDLE_BANNER` – a message that appears as a banner in the middle of the page.
	// - `CAROUSEL` – a scrollable layout of up to five unique messages.
	Layout *string `field:"optional" json:"layout" yaml:"layout"`
}

