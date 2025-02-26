package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// The props for a EMR Containers StartJobRun Task.
//
// Example:
//   tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &EmrContainersStartJobRunProps{
//   	VirtualCluster: tasks.VirtualClusterInput_FromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
//   	ReleaseLabel: tasks.ReleaseLabel_EMR_6_2_0(),
//   	JobName: jsii.String("EMR-Containers-Job"),
//   	JobDriver: &JobDriver{
//   		SparkSubmitJobDriver: &SparkSubmitJobDriver{
//   			EntryPoint: sfn.TaskInput_FromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
//   		},
//   	},
//   	ApplicationConfig: []applicationConfiguration{
//   		&applicationConfiguration{
//   			Classification: tasks.Classification_SPARK_DEFAULTS(),
//   			Properties: map[string]*string{
//   				"spark.executor.instances": jsii.String("1"),
//   				"spark.executor.memory": jsii.String("512M"),
//   			},
//   		},
//   	},
//   })
//
type EmrContainersStartJobRunProps struct {
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
	// Default: - No application config.
	//
	ApplicationConfig *[]*ApplicationConfiguration `field:"optional" json:"applicationConfig" yaml:"applicationConfig"`
	// The execution role for the job run.
	//
	// If `virtualClusterId` is from a JSON input path, an execution role must be provided.
	// If an execution role is provided, follow the documentation to update the role trust policy.
	// See: https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/setting-up-trust-policy.html
	//
	// Default: - Automatically generated only when the provided `virtualClusterId` is not an encoded JSON path.
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of the job run.
	// Default: - No job run name.
	//
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Configuration for monitoring the job run.
	// See: https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_MonitoringConfiguration.html
	//
	// Default: - logging enabled and resources automatically generated if `monitoring.logging` is set to `true`
	//
	Monitoring *Monitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// The tags assigned to job runs.
	// Default: - None.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

