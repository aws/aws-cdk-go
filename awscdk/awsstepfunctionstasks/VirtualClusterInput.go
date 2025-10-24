package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Class that returns a virtual cluster's id depending on input type.
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
//   	ApplicationConfig: []ApplicationConfiguration{
//   		&ApplicationConfiguration{
//   			Classification: tasks.Classification_SPARK_DEFAULTS(),
//   			Properties: map[string]*string{
//   				"spark.executor.instances": jsii.String("1"),
//   				"spark.executor.memory": jsii.String("512M"),
//   			},
//   		},
//   	},
//   })
//
type VirtualClusterInput interface {
	// The VirtualCluster Id.
	Id() *string
}

// The jsii proxy struct for VirtualClusterInput
type jsiiProxy_VirtualClusterInput struct {
	_ byte // padding
}

func (j *jsiiProxy_VirtualClusterInput) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}


// Input for a virtualClusterId from a Task Input.
func VirtualClusterInput_FromTaskInput(taskInput awsstepfunctions.TaskInput) VirtualClusterInput {
	_init_.Initialize()

	if err := validateVirtualClusterInput_FromTaskInputParameters(taskInput); err != nil {
		panic(err)
	}
	var returns VirtualClusterInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.VirtualClusterInput",
		"fromTaskInput",
		[]interface{}{taskInput},
		&returns,
	)

	return returns
}

// Input for virtualClusterId from a literal string.
func VirtualClusterInput_FromVirtualClusterId(virtualClusterId *string) VirtualClusterInput {
	_init_.Initialize()

	if err := validateVirtualClusterInput_FromVirtualClusterIdParameters(virtualClusterId); err != nil {
		panic(err)
	}
	var returns VirtualClusterInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.VirtualClusterInput",
		"fromVirtualClusterId",
		[]interface{}{virtualClusterId},
		&returns,
	)

	return returns
}

