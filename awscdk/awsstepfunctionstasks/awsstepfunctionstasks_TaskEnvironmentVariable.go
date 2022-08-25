package awsstepfunctionstasks


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
	// Use `JsonPath` class's static methods to specify name from a JSON path.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the environment variable.
	//
	// Use `JsonPath` class's static methods to specify value from a JSON path.
	Value *string `field:"required" json:"value" yaml:"value"`
}

