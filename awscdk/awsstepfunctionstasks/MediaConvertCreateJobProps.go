package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for creating a MediaConvert Job.
//
// See the CreateJob API for complete documentation.
//
// Example:
//   tasks.NewMediaConvertCreateJob(this, jsii.String("CreateJob"), &MediaConvertCreateJobProps{
//   	CreateJobRequest: map[string]interface{}{
//   		"Role": jsii.String("arn:aws:iam::123456789012:role/MediaConvertRole"),
//   		"Settings": map[string][]map[string]interface{}{
//   			"OutputGroups": []map[string]interface{}{
//   				map[string]interface{}{
//   					"Outputs": []map[string]interface{}{
//   						map[string]interface{}{
//   							"ContainerSettings": map[string]*string{
//   								"Container": jsii.String("MP4"),
//   							},
//   							"VideoDescription": map[string]map[string]interface{}{
//   								"CodecSettings": map[string]interface{}{
//   									"Codec": jsii.String("H_264"),
//   									"H264Settings": map[string]interface{}{
//   										"MaxBitrate": jsii.Number(1000),
//   										"RateControlMode": jsii.String("QVBR"),
//   										"SceneChangeDetect": jsii.String("TRANSITION_DETECTION"),
//   									},
//   								},
//   							},
//   							"AudioDescriptions": []map[string]map[string]interface{}{
//   								map[string]map[string]interface{}{
//   									"CodecSettings": map[string]interface{}{
//   										"Codec": jsii.String("AAC"),
//   										"AacSettings": map[string]interface{}{
//   											"Bitrate": jsii.Number(96000),
//   											"CodingMode": jsii.String("CODING_MODE_2_0"),
//   											"SampleRate": jsii.Number(48000),
//   										},
//   									},
//   								},
//   							},
//   						},
//   					},
//   					"OutputGroupSettings": map[string]interface{}{
//   						"Type": jsii.String("FILE_GROUP_SETTINGS"),
//   						"FileGroupSettings": map[string]*string{
//   							"Destination": jsii.String("s3://EXAMPLE-DESTINATION-BUCKET/"),
//   						},
//   					},
//   				},
//   			},
//   			"Inputs": []map[string]interface{}{
//   				map[string]interface{}{
//   					"AudioSelectors": map[string]map[string]*string{
//   						"Audio Selector 1": map[string]*string{
//   							"DefaultSelection": jsii.String("DEFAULT"),
//   						},
//   					},
//   					"FileInput": jsii.String("s3://EXAMPLE-SOURCE-BUCKET/EXAMPLE-SOURCE_FILE"),
//   				},
//   			},
//   		},
//   	},
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   })
//
// See: https://docs.aws.amazon.com/mediaconvert/latest/apireference/jobs.html#jobspost
//
type MediaConvertCreateJobProps struct {
	// An optional description for this state.
	// Default: - No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
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
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: - The entire task input (JSON path '$').
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
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
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: - The entire JSON node determined by the state input, the task result,
	// and resultPath is passed to the next state (JSON path '$').
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: - Replaces the entire input with the result (JSON path '$').
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
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
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
	// The input data for the MediaConvert Create Job invocation.
	CreateJobRequest *map[string]interface{} `field:"required" json:"createJobRequest" yaml:"createJobRequest"`
}

