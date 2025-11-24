package mixinsawsgameliftstreams


// Configuration settings that identify the operating system for an application resource.
//
// This can also include a compatibility layer and other drivers.
//
// A runtime environment can be one of the following:
//
// - For Linux applications
//
// - Ubuntu 22.04 LTS ( `Type=UBUNTU, Version=22_04_LTS` )
// - For Windows applications
//
// - Microsoft Windows Server 2022 Base ( `Type=WINDOWS, Version=2022` )
// - Proton 9.0-2 ( `Type=PROTON, Version=20250516` )
// - Proton 8.0-5 ( `Type=PROTON, Version=20241007` )
// - Proton 8.0-2c ( `Type=PROTON, Version=20230704` )
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   runtimeEnvironmentProperty := &RuntimeEnvironmentProperty{
//   	Type: jsii.String("type"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-application-runtimeenvironment.html
//
type CfnApplicationPropsMixin_RuntimeEnvironmentProperty struct {
	// The operating system and other drivers.
	//
	// For Proton, this also includes the Proton compatibility layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-application-runtimeenvironment.html#cfn-gameliftstreams-application-runtimeenvironment-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Versioned container environment for the application operating system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-application-runtimeenvironment.html#cfn-gameliftstreams-application-runtimeenvironment-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

