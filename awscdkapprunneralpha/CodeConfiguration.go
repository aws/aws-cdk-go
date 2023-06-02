package awscdkapprunneralpha


// Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   var runtime runtime
//   var secret secret
//
//   codeConfiguration := &CodeConfiguration{
//   	ConfigurationSource: apprunner_alpha.ConfigurationSourceType_REPOSITORY,
//
//   	// the properties below are optional
//   	ConfigurationValues: &CodeConfigurationValues{
//   		Runtime: runtime,
//
//   		// the properties below are optional
//   		BuildCommand: jsii.String("buildCommand"),
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		EnvironmentSecrets: map[string]*secret{
//   			"environmentSecretsKey": secret,
//   		},
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		Port: jsii.String("port"),
//   		StartCommand: jsii.String("startCommand"),
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

