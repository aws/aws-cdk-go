package previewawspinpointmixins


// Properties for CfnInAppTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var customConfig interface{}
//   var tags interface{}
//
//   cfnInAppTemplateMixinProps := &CfnInAppTemplateMixinProps{
//   	Content: []interface{}{
//   		&InAppMessageContentProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BodyConfig: &BodyConfigProperty{
//   				Alignment: jsii.String("alignment"),
//   				Body: jsii.String("body"),
//   				TextColor: jsii.String("textColor"),
//   			},
//   			HeaderConfig: &HeaderConfigProperty{
//   				Alignment: jsii.String("alignment"),
//   				Header: jsii.String("header"),
//   				TextColor: jsii.String("textColor"),
//   			},
//   			ImageUrl: jsii.String("imageUrl"),
//   			PrimaryBtn: &ButtonConfigProperty{
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
//   			SecondaryBtn: &ButtonConfigProperty{
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
//   	Tags: tags,
//   	TemplateDescription: jsii.String("templateDescription"),
//   	TemplateName: jsii.String("templateName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-inapptemplate.html
//
type CfnInAppTemplateMixinProps struct {
	// An object that contains information about the content of an in-app message, including its title and body text, text colors, background colors, images, buttons, and behaviors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-inapptemplate.html#cfn-pinpoint-inapptemplate-content
	//
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// Custom data, in the form of key-value pairs, that is included in an in-app messaging payload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-inapptemplate.html#cfn-pinpoint-inapptemplate-customconfig
	//
	CustomConfig interface{} `field:"optional" json:"customConfig" yaml:"customConfig"`
	// A string that determines the appearance of the in-app message. You can specify one of the following:.
	//
	// - `BOTTOM_BANNER` – a message that appears as a banner at the bottom of the page.
	// - `TOP_BANNER` – a message that appears as a banner at the top of the page.
	// - `OVERLAYS` – a message that covers entire screen.
	// - `MOBILE_FEED` – a message that appears in a window in front of the page.
	// - `MIDDLE_BANNER` – a message that appears as a banner in the middle of the page.
	// - `CAROUSEL` – a scrollable layout of up to five unique messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-inapptemplate.html#cfn-pinpoint-inapptemplate-layout
	//
	Layout *string `field:"optional" json:"layout" yaml:"layout"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-inapptemplate.html#cfn-pinpoint-inapptemplate-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// An optional description of the in-app template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-inapptemplate.html#cfn-pinpoint-inapptemplate-templatedescription
	//
	TemplateDescription *string `field:"optional" json:"templateDescription" yaml:"templateDescription"`
	// The name of the in-app message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-inapptemplate.html#cfn-pinpoint-inapptemplate-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

