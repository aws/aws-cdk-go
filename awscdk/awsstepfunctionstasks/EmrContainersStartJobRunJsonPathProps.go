package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for calling EMR Containers StartJobRun using JSONPath.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationConfiguration_ applicationConfiguration
//   var assign interface{}
//   var bucket bucket
//   var classification classification
//   var logGroup logGroup
//   var releaseLabel releaseLabel
//   var resultSelector interface{}
//   var role role
//   var taskInput taskInput
//   var taskRole taskRole
//   var timeout timeout
//   var virtualClusterInput virtualClusterInput
//
//   emrContainersStartJobRunJsonPathProps := &EmrContainersStartJobRunJsonPathProps{
//   	JobDriver: &JobDriver{
//   		SparkSubmitJobDriver: &SparkSubmitJobDriver{
//   			EntryPoint: taskInput,
//
//   			// the properties below are optional
//   			EntryPointArguments: taskInput,
//   			SparkSubmitParameters: jsii.String("sparkSubmitParameters"),
//   		},
//   	},
//   	ReleaseLabel: releaseLabel,
//   	VirtualCluster: virtualClusterInput,
//
//   	// the properties below are optional
//   	ApplicationConfig: []*applicationConfiguration{
//   		&applicationConfiguration{
//   			Classification: classification,
//
//   			// the properties below are optional
//   			NestedConfig: []*applicationConfiguration{
//   				applicationConfiguration_,
//   			},
//   			Properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Comment: jsii.String("comment"),
//   	Credentials: &Credentials{
//   		Role: taskRole,
//   	},
//   	ExecutionRole: role,
//   	Heartbeat: cdk.Duration_Minutes(jsii.Number(30)),
//   	HeartbeatTimeout: timeout,
//   	InputPath: jsii.String("inputPath"),
//   	IntegrationPattern: awscdk.Aws_stepfunctions.IntegrationPattern_REQUEST_RESPONSE,
//   	JobName: jsii.String("jobName"),
//   	Monitoring: &Monitoring{
//   		LogBucket: bucket,
//   		Logging: jsii.Boolean(false),
//   		LogGroup: logGroup,
//   		LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   		PersistentAppUI: jsii.Boolean(false),
//   	},
//   	OutputPath: jsii.String("outputPath"),
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	ResultPath: jsii.String("resultPath"),
//   	ResultSelector: map[string]interface{}{
//   		"resultSelectorKey": resultSelector,
//   	},
//   	StateName: jsii.String("stateName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TaskTimeout: timeout,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type EmrContainersStartJobRunJsonPathProps struct {
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

