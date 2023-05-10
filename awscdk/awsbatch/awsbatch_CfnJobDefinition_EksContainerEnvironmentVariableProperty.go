package awsbatch


// An environment variable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksContainerEnvironmentVariableProperty := &EksContainerEnvironmentVariableProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Value: jsii.String("value"),
//   }
//
type CfnJobDefinition_EksContainerEnvironmentVariableProperty struct {
	// The name of the environment variable.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the environment variable.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

