package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Run an ECS/EC2 Task in a StepFunctions workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var placementConstraint placementConstraint
//   var placementStrategy placementStrategy
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var taskDefinition taskDefinition
//
//   runEcsEc2Task := awscdk.Aws_stepfunctions_tasks.NewRunEcsEc2Task(&runEcsEc2TaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerDefinition: containerDefinition,
//
//   			// the properties below are optional
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			cpu: jsii.Number(123),
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			memoryLimit: jsii.Number(123),
//   			memoryReservation: jsii.Number(123),
//   		},
//   	},
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	placementConstraints: []*placementConstraint{
//   		placementConstraint,
//   	},
//   	placementStrategies: []*placementStrategy{
//   		placementStrategy,
//   	},
//   	securityGroup: securityGroup,
//   	subnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
//   	},
//   })
//
// Deprecated: - replaced by `EcsRunTask`.
type RunEcsEc2Task interface {
	EcsRunTaskBase
	// Manage allowed network traffic for this service.
	// Deprecated: - replaced by `EcsRunTask`.
	Connections() awsec2.Connections
	// Called when the task object is used in a workflow.
	// Deprecated: - replaced by `EcsRunTask`.
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
	// Deprecated: - replaced by `EcsRunTask`.
	ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup)
}

// The jsii proxy struct for RunEcsEc2Task
type jsiiProxy_RunEcsEc2Task struct {
	jsiiProxy_EcsRunTaskBase
}

func (j *jsiiProxy_RunEcsEc2Task) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


// Deprecated: - replaced by `EcsRunTask`.
func NewRunEcsEc2Task(props *RunEcsEc2TaskProps) RunEcsEc2Task {
	_init_.Initialize()

	if err := validateNewRunEcsEc2TaskParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RunEcsEc2Task{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunEcsEc2Task",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: - replaced by `EcsRunTask`.
func NewRunEcsEc2Task_Override(r RunEcsEc2Task, props *RunEcsEc2TaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunEcsEc2Task",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_RunEcsEc2Task) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	if err := r.validateBindParameters(task); err != nil {
		panic(err)
	}
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunEcsEc2Task) ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup) {
	if err := r.validateConfigureAwsVpcNetworkingParameters(vpc, subnetSelection); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"configureAwsVpcNetworking",
		[]interface{}{vpc, assignPublicIp, subnetSelection, securityGroup},
	)
}

