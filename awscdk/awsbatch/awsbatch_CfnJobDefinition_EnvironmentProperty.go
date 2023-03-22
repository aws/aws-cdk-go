package awsbatch


// The Environment property type specifies environment variables to use in a job definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentProperty := &environmentProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnJobDefinition_EnvironmentProperty struct {
	// The name of the environment variable.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the environment variable.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

