package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Amazon EMR release version to use for the job run.
//
// Can be extended to include new EMR releases
//
// For example, `new ReleaseLabel('emr-x.xx.x-latest');`
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
type ReleaseLabel interface {
	// A literal string that contains the release-version ex.
	//
	// 'emr-x.x.x-latest'
	Label() *string
}

// The jsii proxy struct for ReleaseLabel
type jsiiProxy_ReleaseLabel struct {
	_ byte // padding
}

func (j *jsiiProxy_ReleaseLabel) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}


// Initializes the label string.
func NewReleaseLabel(label *string) ReleaseLabel {
	_init_.Initialize()

	if err := validateNewReleaseLabelParameters(label); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReleaseLabel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.ReleaseLabel",
		[]interface{}{label},
		&j,
	)

	return &j
}

// Initializes the label string.
func NewReleaseLabel_Override(r ReleaseLabel, label *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.ReleaseLabel",
		[]interface{}{label},
		r,
	)
}

func ReleaseLabel_EMR_5_32_0() ReleaseLabel {
	_init_.Initialize()
	var returns ReleaseLabel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.ReleaseLabel",
		"EMR_5_32_0",
		&returns,
	)
	return returns
}

func ReleaseLabel_EMR_5_33_0() ReleaseLabel {
	_init_.Initialize()
	var returns ReleaseLabel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.ReleaseLabel",
		"EMR_5_33_0",
		&returns,
	)
	return returns
}

func ReleaseLabel_EMR_6_2_0() ReleaseLabel {
	_init_.Initialize()
	var returns ReleaseLabel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.ReleaseLabel",
		"EMR_6_2_0",
		&returns,
	)
	return returns
}

func ReleaseLabel_EMR_6_3_0() ReleaseLabel {
	_init_.Initialize()
	var returns ReleaseLabel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.ReleaseLabel",
		"EMR_6_3_0",
		&returns,
	)
	return returns
}

