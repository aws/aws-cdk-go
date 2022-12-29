// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha


// Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   var runtime runtime
//
//   codeConfiguration := &codeConfiguration{
//   	configurationSource: apprunner_alpha.configurationSourceType_REPOSITORY,
//
//   	// the properties below are optional
//   	configurationValues: &codeConfigurationValues{
//   		runtime: runtime,
//
//   		// the properties below are optional
//   		buildCommand: jsii.String("buildCommand"),
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		port: jsii.String("port"),
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfiguration.html
//
// Experimental.
type CodeConfiguration struct {
	// The source of the App Runner configuration.
	// Experimental.
	ConfigurationSource ConfigurationSourceType `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a apprunner.yaml file in the
	// source code repository (or ignoring the file if it exists).
	// Experimental.
	ConfigurationValues *CodeConfigurationValues `field:"optional" json:"configurationValues" yaml:"configurationValues"`
}

