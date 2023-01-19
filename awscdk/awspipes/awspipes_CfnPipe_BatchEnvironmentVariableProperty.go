package awspipes


// The environment variables to send to the container.
//
// You can add new environment variables, which are added to the container at launch, or you can override the existing environment variables from the Docker image or the task definition.
//
// > Environment variables cannot start with " `AWS Batch` ". This naming convention is reserved for variables that AWS Batch sets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchEnvironmentVariableProperty := &batchEnvironmentVariableProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnPipe_BatchEnvironmentVariableProperty struct {
	// The name of the key-value pair.
	//
	// For environment variables, this is the name of the environment variable.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the key-value pair.
	//
	// For environment variables, this is the value of the environment variable.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

