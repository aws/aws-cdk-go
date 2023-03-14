package awspinpoint


// Specifies the configuration of an in-app message, including its header, body, buttons, colors, and images.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageContentProperty := &inAppMessageContentProperty{
//   	backgroundColor: jsii.String("backgroundColor"),
//   	bodyConfig: &bodyConfigProperty{
//   		alignment: jsii.String("alignment"),
//   		body: jsii.String("body"),
//   		textColor: jsii.String("textColor"),
//   	},
//   	headerConfig: &headerConfigProperty{
//   		alignment: jsii.String("alignment"),
//   		header: jsii.String("header"),
//   		textColor: jsii.String("textColor"),
//   	},
//   	imageUrl: jsii.String("imageUrl"),
//   	primaryBtn: &buttonConfigProperty{
//   		android: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		defaultConfig: &defaultButtonConfigurationProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			borderRadius: jsii.Number(123),
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   			text: jsii.String("text"),
//   			textColor: jsii.String("textColor"),
//   		},
//   		ios: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		web: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   	},
//   	secondaryBtn: &buttonConfigProperty{
//   		android: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		defaultConfig: &defaultButtonConfigurationProperty{
//   			backgroundColor: jsii.String("backgroundColor"),
//   			borderRadius: jsii.Number(123),
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   			text: jsii.String("text"),
//   			textColor: jsii.String("textColor"),
//   		},
//   		ios: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   		web: &overrideButtonConfigurationProperty{
//   			buttonAction: jsii.String("buttonAction"),
//   			link: jsii.String("link"),
//   		},
//   	},
//   }
//
type CfnInAppTemplate_InAppMessageContentProperty struct {
	// The background color for an in-app message banner, expressed as a hex color code (such as #000000 for black).
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// An object that contains configuration information about the header or title text of the in-app message.
	BodyConfig interface{} `field:"optional" json:"bodyConfig" yaml:"bodyConfig"`
	// An object that contains configuration information about the header or title text of the in-app message.
	HeaderConfig interface{} `field:"optional" json:"headerConfig" yaml:"headerConfig"`
	// The URL of the image that appears on an in-app message banner.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// An object that contains configuration information about the primary button in an in-app message.
	PrimaryBtn interface{} `field:"optional" json:"primaryBtn" yaml:"primaryBtn"`
	// An object that contains configuration information about the secondary button in an in-app message.
	SecondaryBtn interface{} `field:"optional" json:"secondaryBtn" yaml:"secondaryBtn"`
}

