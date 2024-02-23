package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Defines an action for an extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventDestination iEventDestination
//   var role role
//
//   action := appconfig_alpha.NewAction(&ActionProps{
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
//   })
//
// Deprecated.
type Action interface {
	// The action points that will trigger the extension action.
	// Deprecated.
	ActionPoints() *[]ActionPoint
	// The description for the action.
	// Deprecated.
	Description() *string
	// The event destination for the action.
	// Deprecated.
	EventDestination() IEventDestination
	// The execution role for the action.
	// Deprecated.
	ExecutionRole() awsiam.IRole
	// The flag that specifies whether to create the execution role.
	// Deprecated.
	InvokeWithoutExecutionRole() *bool
	// The name for the action.
	// Deprecated.
	Name() *string
}

// The jsii proxy struct for Action
type jsiiProxy_Action struct {
	_ byte // padding
}

func (j *jsiiProxy_Action) ActionPoints() *[]ActionPoint {
	var returns *[]ActionPoint
	_jsii_.Get(
		j,
		"actionPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Action) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Action) EventDestination() IEventDestination {
	var returns IEventDestination
	_jsii_.Get(
		j,
		"eventDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Action) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Action) InvokeWithoutExecutionRole() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"invokeWithoutExecutionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Action) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Deprecated.
func NewAction(props *ActionProps) Action {
	_init_.Initialize()

	if err := validateNewActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Action{}

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.Action",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated.
func NewAction_Override(a Action, props *ActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.Action",
		[]interface{}{props},
		a,
	)
}

