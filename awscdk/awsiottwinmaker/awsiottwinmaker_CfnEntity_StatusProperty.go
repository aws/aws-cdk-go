package awsiottwinmaker


// The current status of the entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var error interface{}
//
//   statusProperty := &statusProperty{
//   	error: error,
//   	state: jsii.String("state"),
//   }
//
type CfnEntity_StatusProperty struct {
	// The error message.
	Error interface{} `field:"optional" json:"error" yaml:"error"`
	// The current state of the entity, component, component type, or workspace.
	//
	// Valid Values: `CREATING | UPDATING | DELETING | ACTIVE | ERROR`.
	State *string `field:"optional" json:"state" yaml:"state"`
}

