package awspinpoint


// Specifies the configuration of a button that appears in an in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageButtonProperty := &inAppMessageButtonProperty{
//   	android: &overrideButtonConfigurationProperty{
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   	},
//   	defaultConfig: &defaultButtonConfigurationProperty{
//   		backgroundColor: jsii.String("backgroundColor"),
//   		borderRadius: jsii.Number(123),
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   		text: jsii.String("text"),
//   		textColor: jsii.String("textColor"),
//   	},
//   	ios: &overrideButtonConfigurationProperty{
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   	},
//   	web: &overrideButtonConfigurationProperty{
//   		buttonAction: jsii.String("buttonAction"),
//   		link: jsii.String("link"),
//   	},
//   }
//
type CfnCampaign_InAppMessageButtonProperty struct {
	// An object that defines the default behavior for a button in in-app messages sent to Android.
	Android interface{} `field:"optional" json:"android" yaml:"android"`
	// An object that defines the default behavior for a button in an in-app message.
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// An object that defines the default behavior for a button in in-app messages sent to iOS devices.
	Ios interface{} `field:"optional" json:"ios" yaml:"ios"`
	// An object that defines the default behavior for a button in in-app messages for web applications.
	Web interface{} `field:"optional" json:"web" yaml:"web"`
}

