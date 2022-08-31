package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Amazon EMR release version to use for the job run.
//
// Can be extended to include new EMR releases
//
// For example, `new ReleaseLabel('emr-x.xx.x-latest');`
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
// Experimental.
type ReleaseLabel interface {
	// A literal string that contains the release-version ex.
	//
	// 'emr-x.x.x-latest'
	// Experimental.
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
// Experimental.
func NewReleaseLabel(label *string) ReleaseLabel {
	_init_.Initialize()

	j := jsiiProxy_ReleaseLabel{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.ReleaseLabel",
		[]interface{}{label},
		&j,
	)

	return &j
}

// Initializes the label string.
// Experimental.
func NewReleaseLabel_Override(r ReleaseLabel, label *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.ReleaseLabel",
		[]interface{}{label},
		r,
	)
}

func ReleaseLabel_EMR_5_32_0() ReleaseLabel {
	_init_.Initialize()
	var returns ReleaseLabel
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.ReleaseLabel",
		"EMR_5_32_0",
		&returns,
	)
	return returns
}

func ReleaseLabel_EMR_5_33_0() ReleaseLabel {
	_init_.Initialize()
	var returns ReleaseLabel
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.ReleaseLabel",
		"EMR_5_33_0",
		&returns,
	)
	return returns
}

func ReleaseLabel_EMR_6_2_0() ReleaseLabel {
	_init_.Initialize()
	var returns ReleaseLabel
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.ReleaseLabel",
		"EMR_6_2_0",
		&returns,
	)
	return returns
}

func ReleaseLabel_EMR_6_3_0() ReleaseLabel {
	_init_.Initialize()
	var returns ReleaseLabel
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.ReleaseLabel",
		"EMR_6_3_0",
		&returns,
	)
	return returns
}

