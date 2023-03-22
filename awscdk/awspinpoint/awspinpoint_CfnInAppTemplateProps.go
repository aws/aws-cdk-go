package awspinpoint


// Properties for defining a `CfnInAppTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customConfig interface{}
//   var tags interface{}
//
//   cfnInAppTemplateProps := &cfnInAppTemplateProps{
//   	templateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	content: []interface{}{
//   		&inAppMessageContentProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			bodyConfig: &bodyConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				body: jsii.String("body"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			headerConfig: &headerConfigProperty{
//   				alignment: jsii.String("alignment"),
//   				header: jsii.String("header"),
//   				textColor: jsii.String("textColor"),
//   			},
//   			imageUrl: jsii.String("imageUrl"),
//   			primaryBtn: &buttonConfigProperty{
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
//   			secondaryBtn: &buttonConfigProperty{
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
//   	tags: tags,
//   	templateDescription: jsii.String("templateDescription"),
//   }
//
type CfnInAppTemplateProps struct {
	// The name of the in-app message template.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// An object that contains information about the content of an in-app message, including its title and body text, text colors, background colors, images, buttons, and behaviors.
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// Custom data, in the form of key-value pairs, that is included in an in-app messaging payload.
	CustomConfig interface{} `field:"optional" json:"customConfig" yaml:"customConfig"`
	// A string that determines the appearance of the in-app message. You can specify one of the following:.
	//
	// - `BOTTOM_BANNER` – a message that appears as a banner at the bottom of the page.
	// - `TOP_BANNER` – a message that appears as a banner at the top of the page.
	// - `OVERLAYS` – a message that covers entire screen.
	// - `MOBILE_FEED` – a message that appears in a window in front of the page.
	// - `MIDDLE_BANNER` – a message that appears as a banner in the middle of the page.
	// - `CAROUSEL` – a scrollable layout of up to five unique messages.
	Layout *string `field:"optional" json:"layout" yaml:"layout"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// An optional description of the in-app template.
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
}

