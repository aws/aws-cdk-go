package awspinpoint


// Specifies the behavior of buttons that appear in an in-app message template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buttonConfigProperty := &buttonConfigProperty{
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
type CfnInAppTemplate_ButtonConfigProperty struct {
	// Optional button configuration to use for in-app messages sent to Android devices.
	//
	// This button configuration overrides the default button configuration.
	Android interface{} `field:"optional" json:"android" yaml:"android"`
	// Specifies the default behavior of a button that appears in an in-app message.
	//
	// You can optionally add button configurations that specifically apply to iOS, Android, or web browser users.
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// Optional button configuration to use for in-app messages sent to iOS devices.
	//
	// This button configuration overrides the default button configuration.
	Ios interface{} `field:"optional" json:"ios" yaml:"ios"`
	// Optional button configuration to use for in-app messages sent to web applications.
	//
	// This button configuration overrides the default button configuration.
	Web interface{} `field:"optional" json:"web" yaml:"web"`
}

