package awseventstargets


// An environment variable to be set in the container run as a task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskEnvironmentVariable := &taskEnvironmentVariable{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type TaskEnvironmentVariable struct {
	// Name for the environment variable.
	//
	// Exactly one of `name` and `namePath` must be specified.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the environment variable.
	//
	// Exactly one of `value` and `valuePath` must be specified.
	Value *string `field:"required" json:"value" yaml:"value"`
}

