package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
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
// Experimental.
type EmrContainersStartJobRunProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The job driver for the job run.
	// See: https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_JobDriver.html
	//
	// Experimental.
	JobDriver *JobDriver `field:"required" json:"jobDriver" yaml:"jobDriver"`
	// The Amazon EMR release version to use for the job run.
	// Experimental.
	ReleaseLabel ReleaseLabel `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// The ID of the virtual cluster where the job will be run.
	// Experimental.
	VirtualCluster VirtualClusterInput `field:"required" json:"virtualCluster" yaml:"virtualCluster"`
	// The configurations for the application running in the job run.
	//
	// Maximum of 100 items.
	// See: https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_Configuration.html
	//
	// Experimental.
	ApplicationConfig *[]*ApplicationConfiguration `field:"optional" json:"applicationConfig" yaml:"applicationConfig"`
	// The execution role for the job run.
	//
	// If `virtualClusterId` is from a JSON input path, an execution role must be provided.
	// If an execution role is provided, follow the documentation to update the role trust policy.
	// See: https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/setting-up-trust-policy.html
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of the job run.
	// Experimental.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Configuration for monitoring the job run.
	// See: https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_MonitoringConfiguration.html
	//
	// Experimental.
	Monitoring *Monitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// The tags assigned to job runs.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

