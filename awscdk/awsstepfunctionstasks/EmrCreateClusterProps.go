package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for EmrCreateCluster.
//
// See the RunJobFlow API for complete documentation on input parameters.
//
// Example:
//   clusterRole := iam.NewRole(this, jsii.String("ClusterRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
//   })
//
//   serviceRole := iam.NewRole(this, jsii.String("ServiceRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("elasticmapreduce.amazonaws.com")),
//   })
//
//   autoScalingRole := iam.NewRole(this, jsii.String("AutoScalingRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("elasticmapreduce.amazonaws.com")),
//   })
//
//   autoScalingRole.AssumeRolePolicy.AddStatements(
//   iam.NewPolicyStatement(&PolicyStatementProps{
//   	Effect: iam.Effect_ALLOW,
//   	Principals: []iPrincipal{
//   		iam.NewServicePrincipal(jsii.String("application-autoscaling.amazonaws.com")),
//   	},
//   	Actions: []*string{
//   		jsii.String("sts:AssumeRole"),
//   	},
//   }))
//
//   tasks.NewEmrCreateCluster(this, jsii.String("Create Cluster"), &EmrCreateClusterProps{
//   	Instances: &InstancesConfigProperty{
//   	},
//   	ClusterRole: ClusterRole,
//   	Name: sfn.TaskInput_FromJsonPathAt(jsii.String("$.ClusterName")).value,
//   	ServiceRole: ServiceRole,
//   	AutoScalingRole: AutoScalingRole,
//   })
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_RunJobFlow.html
//
type EmrCreateClusterProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// A specification of the number and type of Amazon EC2 instances.
	Instances *EmrCreateCluster_InstancesConfigProperty `field:"required" json:"instances" yaml:"instances"`
	// The Name of the Cluster.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A JSON string for selecting additional features.
	AdditionalInfo *string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// A case-insensitive list of applications for Amazon EMR to install and configure when launching the cluster.
	Applications *[]*EmrCreateCluster_ApplicationConfigProperty `field:"optional" json:"applications" yaml:"applications"`
	// An IAM role for automatic scaling policies.
	AutoScalingRole awsiam.IRole `field:"optional" json:"autoScalingRole" yaml:"autoScalingRole"`
	// A list of bootstrap actions to run before Hadoop starts on the cluster nodes.
	BootstrapActions *[]*EmrCreateCluster_BootstrapActionConfigProperty `field:"optional" json:"bootstrapActions" yaml:"bootstrapActions"`
	// Also called instance profile and EC2 role.
	//
	// An IAM role for an EMR cluster. The EC2 instances of the cluster assume this role.
	//
	// This attribute has been renamed from jobFlowRole to clusterRole to align with other ERM/StepFunction integration parameters.
	ClusterRole awsiam.IRole `field:"optional" json:"clusterRole" yaml:"clusterRole"`
	// The list of configurations supplied for the EMR cluster you are creating.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `field:"optional" json:"configurations" yaml:"configurations"`
	// The ID of a custom Amazon EBS-backed Linux AMI.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// The size of the EBS root device volume of the Linux AMI that is used for each EC2 instance.
	EbsRootVolumeSize awscdk.Size `field:"optional" json:"ebsRootVolumeSize" yaml:"ebsRootVolumeSize"`
	// Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.
	KerberosAttributes *EmrCreateCluster_KerberosAttributesProperty `field:"optional" json:"kerberosAttributes" yaml:"kerberosAttributes"`
	// The location in Amazon S3 to write the log files of the job flow.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.
	ReleaseLabel *string `field:"optional" json:"releaseLabel" yaml:"releaseLabel"`
	// Specifies the way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.
	ScaleDownBehavior EmrCreateCluster_EmrClusterScaleDownBehavior `field:"optional" json:"scaleDownBehavior" yaml:"scaleDownBehavior"`
	// The name of a security configuration to apply to the cluster.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// Specifies the step concurrency level to allow multiple steps to run in parallel.
	//
	// Requires EMR release label 5.28.0 or above.
	// Must be in range [1, 256].
	StepConcurrencyLevel *float64 `field:"optional" json:"stepConcurrencyLevel" yaml:"stepConcurrencyLevel"`
	// A list of tags to associate with a cluster and propagate to Amazon EC2 instances.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// A value of true indicates that all IAM users in the AWS account can perform cluster actions if they have the proper IAM policy permissions.
	VisibleToAllUsers *bool `field:"optional" json:"visibleToAllUsers" yaml:"visibleToAllUsers"`
}

