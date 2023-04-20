package awsapprunner


// Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigurationProperty := &ImageConfigurationProperty{
//   	Port: jsii.String("port"),
//   	RuntimeEnvironmentSecrets: []interface{}{
//   		&KeyValuePairProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	RuntimeEnvironmentVariables: []interface{}{
//   		&KeyValuePairProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	StartCommand: jsii.String("startCommand"),
//   }
//
type CfnService_ImageConfigurationProperty struct {
	// The port that your application listens to in the container.
	//
	// Default: `8080`.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// `CfnService.ImageConfigurationProperty.RuntimeEnvironmentSecrets`.
	RuntimeEnvironmentSecrets interface{} `field:"optional" json:"runtimeEnvironmentSecrets" yaml:"runtimeEnvironmentSecrets"`
	// Environment variables that are available to your running App Runner service.
	//
	// An array of key-value pairs.
	RuntimeEnvironmentVariables interface{} `field:"optional" json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// An optional command that App Runner runs to start the application in the source image.
	//
	// If specified, this command overrides the Docker image’s default start command.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

