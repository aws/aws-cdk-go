package awspinpoint


// Specifies the configuration of a button that appears in an in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageButtonProperty := &InAppMessageButtonProperty{
//   	Android: &OverrideButtonConfigurationProperty{
//   		ButtonAction: jsii.String("buttonAction"),
//   		Link: jsii.String("link"),
//   	},
//   	DefaultConfig: &DefaultButtonConfigurationProperty{
//   		BackgroundColor: jsii.String("backgroundColor"),
//   		BorderRadius: jsii.Number(123),
//   		ButtonAction: jsii.String("buttonAction"),
//   		Link: jsii.String("link"),
//   		Text: jsii.String("text"),
//   		TextColor: jsii.String("textColor"),
//   	},
//   	Ios: &OverrideButtonConfigurationProperty{
//   		ButtonAction: jsii.String("buttonAction"),
//   		Link: jsii.String("link"),
//   	},
//   	Web: &OverrideButtonConfigurationProperty{
//   		ButtonAction: jsii.String("buttonAction"),
//   		Link: jsii.String("link"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagebutton.html
//
type CfnCampaign_InAppMessageButtonProperty struct {
	// An object that defines the default behavior for a button in in-app messages sent to Android.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagebutton.html#cfn-pinpoint-campaign-inappmessagebutton-android
	//
	Android interface{} `field:"optional" json:"android" yaml:"android"`
	// An object that defines the default behavior for a button in an in-app message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagebutton.html#cfn-pinpoint-campaign-inappmessagebutton-defaultconfig
	//
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// An object that defines the default behavior for a button in in-app messages sent to iOS devices.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagebutton.html#cfn-pinpoint-campaign-inappmessagebutton-ios
	//
	Ios interface{} `field:"optional" json:"ios" yaml:"ios"`
	// An object that defines the default behavior for a button in in-app messages for web applications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-inappmessagebutton.html#cfn-pinpoint-campaign-inappmessagebutton-web
	//
	Web interface{} `field:"optional" json:"web" yaml:"web"`
}

