package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Defines an action for an extension.
//
// Example:
//   var fn Function
//
//
//   appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
//   	Actions: []Action{
//   		appconfig.NewAction(&ActionProps{
//   			ActionPoints: []ActionPoint{
//   				appconfig.ActionPoint_ON_DEPLOYMENT_START,
//   			},
//   			EventDestination: appconfig.NewLambdaDestination(fn),
//   		}),
//   	},
//   })
//
type Action interface {
	// The action points that will trigger the extension action.
	ActionPoints() *[]ActionPoint
	// The description for the action.
	Description() *string
	// The event destination for the action.
	EventDestination() IEventDestination
	// The execution role for the action.
	ExecutionRole() awsiam.IRole
	// The flag that specifies whether to create the execution role.
	InvokeWithoutExecutionRole() *bool
	// The name for the action.
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


func NewAction(props *ActionProps) Action {
	_init_.Initialize()

	if err := validateNewActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Action{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.Action",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAction_Override(a Action, props *ActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.Action",
		[]interface{}{props},
		a,
	)
}

