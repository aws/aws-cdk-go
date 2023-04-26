package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// A StepFunctions Task to run a Task on ECS or Fargate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var parameters interface{}
//   var taskDefinition taskDefinition
//
//   ecsRunTaskBase := awscdk.Aws_stepfunctions_tasks.NewEcsRunTaskBase(&EcsRunTaskBaseProps{
//   	Cluster: cluster,
//   	TaskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	ContainerOverrides: []containerOverride{
//   		&containerOverride{
//   			ContainerDefinition: containerDefinition,
//
//   			// the properties below are optional
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: jsii.Number(123),
//   			Environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MemoryLimit: jsii.Number(123),
//   			MemoryReservation: jsii.Number(123),
//   		},
//   	},
//   	IntegrationPattern: awscdk.Aws_stepfunctions.ServiceIntegrationPattern_FIRE_AND_FORGET,
//   	Parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   })
//
// Deprecated: No replacement.
type EcsRunTaskBase interface {
	awsec2.IConnectable
	awsstepfunctions.IStepFunctionsTask
	// Manage allowed network traffic for this service.
	// Deprecated: No replacement.
	Connections() awsec2.Connections
	// Called when the task object is used in a workflow.
	// Deprecated: No replacement.
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
	// Deprecated: No replacement.
	ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup)
}

// The jsii proxy struct for EcsRunTaskBase
type jsiiProxy_EcsRunTaskBase struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

func (j *jsiiProxy_EcsRunTaskBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


// Deprecated: No replacement.
func NewEcsRunTaskBase(props *EcsRunTaskBaseProps) EcsRunTaskBase {
	_init_.Initialize()

	if err := validateNewEcsRunTaskBaseParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsRunTaskBase{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsRunTaskBase",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: No replacement.
func NewEcsRunTaskBase_Override(e EcsRunTaskBase, props *EcsRunTaskBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsRunTaskBase",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EcsRunTaskBase) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	if err := e.validateBindParameters(task); err != nil {
		panic(err)
	}
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsRunTaskBase) ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup) {
	if err := e.validateConfigureAwsVpcNetworkingParameters(vpc, subnetSelection); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"configureAwsVpcNetworking",
		[]interface{}{vpc, assignPublicIp, subnetSelection, securityGroup},
	)
}

