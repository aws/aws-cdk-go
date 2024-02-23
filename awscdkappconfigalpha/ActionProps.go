package awscdkappconfigalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for the Action construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventDestination iEventDestination
//   var role role
//
//   actionProps := &ActionProps{
//   	ActionPoints: []actionPoint{
//   		appconfig_alpha.*actionPoint_PRE_CREATE_HOSTED_CONFIGURATION_VERSION,
//   	},
//   	EventDestination: eventDestination,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExecutionRole: role,
//   	InvokeWithoutExecutionRole: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   }
//
// Deprecated.
type ActionProps struct {
	// The action points that will trigger the extension action.
	// Deprecated.
	ActionPoints *[]ActionPoint `field:"required" json:"actionPoints" yaml:"actionPoints"`
	// The event destination for the action.
	// Deprecated.
	EventDestination IEventDestination `field:"required" json:"eventDestination" yaml:"eventDestination"`
	// The description for the action.
	// Default: - No description.
	//
	// Deprecated.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The execution role for the action.
	// Default: - A role is generated.
	//
	// Deprecated.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The flag that specifies whether or not to create the execution role.
	//
	// If set to true, then the role will not be auto-generated under the assumption
	// there is already the corresponding resource-based policy attached to the event
	// destination. If false, the execution role will be generated if not provided.
	// Default: false.
	//
	// Deprecated.
	InvokeWithoutExecutionRole *bool `field:"optional" json:"invokeWithoutExecutionRole" yaml:"invokeWithoutExecutionRole"`
	// The name for the action.
	// Default: - A name is generated.
	//
	// Deprecated.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

