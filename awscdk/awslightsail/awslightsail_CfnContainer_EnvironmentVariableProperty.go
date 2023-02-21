package awslightsail


// `EnvironmentVariable` is a property of the [Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html) property. It describes the environment variables of a container on a container service which are key-value parameters that provide dynamic configuration of the application or script run by the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentVariableProperty := &environmentVariableProperty{
//   	value: jsii.String("value"),
//   	variable: jsii.String("variable"),
//   }
//
type CfnContainer_EnvironmentVariableProperty struct {
	// The environment variable value.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// The environment variable key.
	Variable *string `field:"optional" json:"variable" yaml:"variable"`
}

