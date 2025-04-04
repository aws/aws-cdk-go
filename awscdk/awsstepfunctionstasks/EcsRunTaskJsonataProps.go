package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for ECS Tasks using JSONata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var ecsLaunchTarget iEcsLaunchTarget
//   var outputs interface{}
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var taskDefinition taskDefinition
//   var taskRole taskRole
//   var timeout timeout
//
//   ecsRunTaskJsonataProps := &EcsRunTaskJsonataProps{
//   	Cluster: cluster,
//   	LaunchTarget: ecsLaunchTarget,
//   	TaskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	AssignPublicIp: jsii.Boolean(false),
//   	Comment: jsii.String("comment"),
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
//   	Cpu: jsii.String("cpu"),
//   	Credentials: &Credentials{
//   		Role: taskRole,
//   	},
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	Heartbeat: cdk.Duration_Minutes(jsii.Number(30)),
//   	HeartbeatTimeout: timeout,
//   	IntegrationPattern: awscdk.Aws_stepfunctions.IntegrationPattern_REQUEST_RESPONSE,
//   	MemoryMiB: jsii.String("memoryMiB"),
//   	Outputs: outputs,
//   	PropagatedTagSource: awscdk.Aws_ecs.PropagatedTagSource_SERVICE,
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	RevisionNumber: jsii.Number(123),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	StateName: jsii.String("stateName"),
//   	Subnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   	TaskTimeout: timeout,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type EcsRunTaskJsonataProps struct {
	// A comment describing this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The name of the query language used by the state.
	//
	// If the state does not contain a `queryLanguage` field,
	// then it will use the query language specified in the top-level `queryLanguage` field.
	// Default: - JSONPath.
	//
	QueryLanguage awsstepfunctions.QueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	// Default: - None (Task is executed using the State Machine's execution role).
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Default: - None.
	//
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	//
	// Depending on the AWS Service, the Service Integration Pattern availability will vary.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-supported-services.html
	//
	// Default: - `IntegrationPattern.REQUEST_RESPONSE` for most tasks.
	// `IntegrationPattern.RUN_JOB` for the following exceptions:
	// `BatchSubmitJob`, `EmrAddStep`, `EmrCreateCluster`, `EmrTerminationCluster`, and `EmrContainersStartJobRun`.
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Default: - None.
	//
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// Used to specify and transform output from the state.
	//
	// When specified, the value overrides the state output default.
	// The output field accepts any JSON value (object, array, string, number, boolean, null).
	// Any string value, including those inside objects or arrays,
	// will be evaluated as JSONata if surrounded by {% %} characters.
	// Output also accepts a JSONata expression directly.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html
	//
	// Default: - $states.result or $states.errorOutput
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// The ECS cluster to run the task on.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// An Amazon ECS launch type determines the type of infrastructure on which your tasks and services are hosted.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html
	//
	LaunchTarget IEcsLaunchTarget `field:"required" json:"launchTarget" yaml:"launchTarget"`
	// [disable-awslint:ref-via-interface] Task Definition used for running tasks in the service.
	//
	// Note: this must be TaskDefinition, and not ITaskDefinition,
	// as it requires properties that are not known for imported task definitions
	// If you want to run a RunTask with an imported task definition,
	// consider using CustomState.
	TaskDefinition awsecs.TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Assign public IP addresses to each task.
	// Default: false.
	//
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Container setting overrides.
	//
	// Specify the container to use and the overrides to apply.
	// Default: - No overrides.
	//
	ContainerOverrides *[]*ContainerOverride `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// Cpu setting override.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskOverride.html
	//
	// Default: - No override.
	//
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Whether ECS Exec should be enabled.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html#ECS-RunTask-request-enableExecuteCommand
	//
	// Default: false.
	//
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Memory setting override.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskOverride.html
	//
	// Default: - No override.
	//
	MemoryMiB *string `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// An error will be received if you specify the SERVICE option when running a task.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html#ECS-RunTask-request-propagateTags
	//
	// Default: - No tags are propagated.
	//
	PropagatedTagSource awsecs.PropagatedTagSource `field:"optional" json:"propagatedTagSource" yaml:"propagatedTagSource"`
	// The revision number of ECS task definition family.
	// Default: - '$latest'.
	//
	RevisionNumber *float64 `field:"optional" json:"revisionNumber" yaml:"revisionNumber"`
	// Existing security groups to use for the tasks.
	// Default: - A new security group is created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Subnets to place the task's ENIs.
	// Default: - Public subnets if assignPublicIp is set. Private subnets otherwise.
	//
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

