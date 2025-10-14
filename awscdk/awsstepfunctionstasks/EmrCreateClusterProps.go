package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for calling an AWS service's API action from your state machine across regions.
//
// Example:
//   tasks.NewEmrCreateCluster(this, jsii.String("CreateCluster"), &EmrCreateClusterProps{
//   	Instances: &InstancesConfigProperty{
//   		InstanceFleets: []instanceFleetConfigProperty{
//   			&instanceFleetConfigProperty{
//   				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_CORE,
//   				InstanceTypeConfigs: []instanceTypeConfigProperty{
//   					&instanceTypeConfigProperty{
//   						InstanceType: jsii.String("m5.xlarge"),
//   					},
//   				},
//   				TargetOnDemandCapacity: jsii.Number(1),
//   			},
//   			&instanceFleetConfigProperty{
//   				InstanceFleetType: tasks.EmrCreateCluster.InstanceRoleType_MASTER,
//   				InstanceTypeConfigs: []*instanceTypeConfigProperty{
//   					&instanceTypeConfigProperty{
//   						InstanceType: jsii.String("m5.xlarge"),
//   					},
//   				},
//   				TargetOnDemandCapacity: jsii.Number(1),
//   			},
//   		},
//   	},
//   	Name: jsii.String("ClusterName"),
//   	ReleaseLabel: jsii.String("emr-7.9.0"),
//   	ManagedScalingPolicy: &ManagedScalingPolicyProperty{
//   		ComputeLimits: &ManagedScalingComputeLimitsProperty{
//   			UnitType: tasks.EmrCreateCluster.ComputeLimitsUnitType_INSTANCE_FLEET_UNITS,
//   			MaximumCapacityUnits: jsii.Number(4),
//   			MinimumCapacityUnits: jsii.Number(1),
//   			MaximumOnDemandCapacityUnits: jsii.Number(4),
//   			MaximumCoreCapacityUnits: jsii.Number(2),
//   		},
//   	},
//   })
//
type EmrCreateClusterProps struct {
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
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: $.
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: $.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
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
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: $.
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Default: - None.
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// A specification of the number and type of Amazon EC2 instances.
	Instances *EmrCreateCluster_InstancesConfigProperty `field:"required" json:"instances" yaml:"instances"`
	// The Name of the Cluster.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A JSON string for selecting additional features.
	// Default: - None.
	//
	AdditionalInfo *string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// A case-insensitive list of applications for Amazon EMR to install and configure when launching the cluster.
	// Default: - EMR selected default.
	//
	Applications *[]*EmrCreateCluster_ApplicationConfigProperty `field:"optional" json:"applications" yaml:"applications"`
	// An IAM role for automatic scaling policies.
	// Default: - A role will be created.
	//
	AutoScalingRole awsiam.IRole `field:"optional" json:"autoScalingRole" yaml:"autoScalingRole"`
	// The amount of idle time after which the cluster automatically terminates.
	//
	// You can specify a minimum of 60 seconds and a maximum of 604800 seconds (seven days).
	// Default: - No timeout.
	//
	AutoTerminationPolicyIdleTimeout awscdk.Duration `field:"optional" json:"autoTerminationPolicyIdleTimeout" yaml:"autoTerminationPolicyIdleTimeout"`
	// A list of bootstrap actions to run before Hadoop starts on the cluster nodes.
	// Default: - None.
	//
	BootstrapActions *[]*EmrCreateCluster_BootstrapActionConfigProperty `field:"optional" json:"bootstrapActions" yaml:"bootstrapActions"`
	// Also called instance profile and EC2 role.
	//
	// An IAM role for an EMR cluster. The EC2 instances of the cluster assume this role.
	//
	// This attribute has been renamed from jobFlowRole to clusterRole to align with other ERM/StepFunction integration parameters.
	// Default: - * A Role will be created.
	//
	ClusterRole awsiam.IRole `field:"optional" json:"clusterRole" yaml:"clusterRole"`
	// The list of configurations supplied for the EMR cluster you are creating.
	// Default: - None.
	//
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `field:"optional" json:"configurations" yaml:"configurations"`
	// The ID of a custom Amazon EBS-backed Linux AMI.
	// Default: - None.
	//
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// The IOPS of the EBS root device volume of the Linux AMI that is used for each EC2 instance.
	//
	// Requires EMR release label 6.15.0 or above.
	// Must be in range [3000, 16000].
	// See: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-custom-ami-root-volume-size.html#emr-root-volume-overview
	//
	// Default: - EMR selected default.
	//
	EbsRootVolumeIops *float64 `field:"optional" json:"ebsRootVolumeIops" yaml:"ebsRootVolumeIops"`
	// The size of the EBS root device volume of the Linux AMI that is used for each EC2 instance.
	// Default: - EMR selected default.
	//
	EbsRootVolumeSize awscdk.Size `field:"optional" json:"ebsRootVolumeSize" yaml:"ebsRootVolumeSize"`
	// The throughput, in MiB/s, of the EBS root device volume of the Linux AMI that is used for each EC2 instance.
	//
	// Requires EMR release label 6.15.0 or above.
	// Must be in range [125, 1000].
	// See: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-custom-ami-root-volume-size.html#emr-root-volume-overview
	//
	// Default: - EMR selected default.
	//
	EbsRootVolumeThroughput *float64 `field:"optional" json:"ebsRootVolumeThroughput" yaml:"ebsRootVolumeThroughput"`
	// Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.
	// Default: - None.
	//
	KerberosAttributes *EmrCreateCluster_KerberosAttributesProperty `field:"optional" json:"kerberosAttributes" yaml:"kerberosAttributes"`
	// The location in Amazon S3 to write the log files of the job flow.
	// Default: - None.
	//
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// The specified managed scaling policy for an Amazon EMR cluster.
	// Default: - None.
	//
	ManagedScalingPolicy *EmrCreateCluster_ManagedScalingPolicyProperty `field:"optional" json:"managedScalingPolicy" yaml:"managedScalingPolicy"`
	// The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.
	// Default: - EMR selected default.
	//
	ReleaseLabel *string `field:"optional" json:"releaseLabel" yaml:"releaseLabel"`
	// Specifies the way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.
	// Default: - EMR selected default.
	//
	ScaleDownBehavior EmrCreateCluster_EmrClusterScaleDownBehavior `field:"optional" json:"scaleDownBehavior" yaml:"scaleDownBehavior"`
	// The name of a security configuration to apply to the cluster.
	// Default: - None.
	//
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.
	// Default: - A role will be created that Amazon EMR service can assume.
	//
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// Specifies the step concurrency level to allow multiple steps to run in parallel.
	//
	// Requires EMR release label 5.28.0 or above.
	// Must be in range [1, 256].
	// Default: 1 - no step concurrency allowed.
	//
	StepConcurrencyLevel *float64 `field:"optional" json:"stepConcurrencyLevel" yaml:"stepConcurrencyLevel"`
	// A list of tags to associate with a cluster and propagate to Amazon EC2 instances.
	// Default: - None.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// A value of true indicates that all IAM users in the AWS account can perform cluster actions if they have the proper IAM policy permissions.
	// Default: true.
	//
	VisibleToAllUsers *bool `field:"optional" json:"visibleToAllUsers" yaml:"visibleToAllUsers"`
}

