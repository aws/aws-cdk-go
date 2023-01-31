package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// The props for a EMR Containers StartJobRun Task.
//
// Example:
//   tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
//   	virtualCluster: tasks.virtualClusterInput.fromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
//   	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
//   	jobName: jsii.String("EMR-Containers-Job"),
//   	jobDriver: &jobDriver{
//   		sparkSubmitJobDriver: &sparkSubmitJobDriver{
//   			entryPoint: sfn.taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
//   		},
//   	},
//   	applicationConfig: []applicationConfiguration{
//   		&applicationConfiguration{
//   			classification: tasks.classification_SPARK_DEFAULTS(),
//   			properties: map[string]*string{
//   				"spark.executor.instances": jsii.String("1"),
//   				"spark.executor.memory": jsii.String("512M"),
//   			},
//   		},
//   	},
//   })
//
type EmrContainersStartJobRunProps struct {
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
	// The job driver for the job run.
	// See: https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_JobDriver.html
	//
	JobDriver *JobDriver `field:"required" json:"jobDriver" yaml:"jobDriver"`
	// The Amazon EMR release version to use for the job run.
	ReleaseLabel ReleaseLabel `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// The ID of the virtual cluster where the job will be run.
	VirtualCluster VirtualClusterInput `field:"required" json:"virtualCluster" yaml:"virtualCluster"`
	// The configurations for the application running in the job run.
	//
	// Maximum of 100 items.
	// See: https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_Configuration.html
	//
	ApplicationConfig *[]*ApplicationConfiguration `field:"optional" json:"applicationConfig" yaml:"applicationConfig"`
	// The execution role for the job run.
	//
	// If `virtualClusterId` is from a JSON input path, an execution role must be provided.
	// If an execution role is provided, follow the documentation to update the role trust policy.
	// See: https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/setting-up-trust-policy.html
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of the job run.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Configuration for monitoring the job run.
	// See: https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_MonitoringConfiguration.html
	//
	Monitoring *Monitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// The tags assigned to job runs.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

