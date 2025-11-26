package previewawspinpointmixins


// Specifies the behavior of buttons that appear in an in-app message template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   buttonConfigProperty := &ButtonConfigProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-buttonconfig.html
//
type CfnInAppTemplatePropsMixin_ButtonConfigProperty struct {
	// Optional button configuration to use for in-app messages sent to Android devices.
	//
	// This button configuration overrides the default button configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-buttonconfig.html#cfn-pinpoint-inapptemplate-buttonconfig-android
	//
	Android interface{} `field:"optional" json:"android" yaml:"android"`
	// Specifies the default behavior of a button that appears in an in-app message.
	//
	// You can optionally add button configurations that specifically apply to iOS, Android, or web browser users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-buttonconfig.html#cfn-pinpoint-inapptemplate-buttonconfig-defaultconfig
	//
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// Optional button configuration to use for in-app messages sent to iOS devices.
	//
	// This button configuration overrides the default button configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-buttonconfig.html#cfn-pinpoint-inapptemplate-buttonconfig-ios
	//
	Ios interface{} `field:"optional" json:"ios" yaml:"ios"`
	// Optional button configuration to use for in-app messages sent to web applications.
	//
	// This button configuration overrides the default button configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-buttonconfig.html#cfn-pinpoint-inapptemplate-buttonconfig-web
	//
	Web interface{} `field:"optional" json:"web" yaml:"web"`
}

