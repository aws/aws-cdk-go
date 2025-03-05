package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for the Action construct.
//
// Example:
//   var fn function
//
//
//   appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
//   	Actions: []action{
//   		appconfig.NewAction(&ActionProps{
//   			ActionPoints: []actionPoint{
//   				appconfig.*actionPoint_ON_DEPLOYMENT_START,
//   			},
//   			EventDestination: appconfig.NewLambdaDestination(fn),
//   		}),
//   	},
//   })
//
type ActionProps struct {
	// The action points that will trigger the extension action.
	ActionPoints *[]ActionPoint `field:"required" json:"actionPoints" yaml:"actionPoints"`
	// The event destination for the action.
	EventDestination IEventDestination `field:"required" json:"eventDestination" yaml:"eventDestination"`
	// The description for the action.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The execution role for the action.
	// Default: - A role is generated.
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The flag that specifies whether or not to create the execution role.
	//
	// If set to true, then the role will not be auto-generated under the assumption
	// there is already the corresponding resource-based policy attached to the event
	// destination. If false, the execution role will be generated if not provided.
	// Default: false.
	//
	InvokeWithoutExecutionRole *bool `field:"optional" json:"invokeWithoutExecutionRole" yaml:"invokeWithoutExecutionRole"`
	// The name for the action.
	// Default: - A name is generated.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

