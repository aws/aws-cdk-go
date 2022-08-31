package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Class that returns a virtual cluster's id depending on input type.
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
type VirtualClusterInput interface {
	// The VirtualCluster Id.
	// Experimental.
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
// Experimental.
func VirtualClusterInput_FromTaskInput(taskInput awsstepfunctions.TaskInput) VirtualClusterInput {
	_init_.Initialize()

	if err := validateVirtualClusterInput_FromTaskInputParameters(taskInput); err != nil {
		panic(err)
	}
	var returns VirtualClusterInput

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.VirtualClusterInput",
		"fromTaskInput",
		[]interface{}{taskInput},
		&returns,
	)

	return returns
}

// Input for virtualClusterId from a literal string.
// Experimental.
func VirtualClusterInput_FromVirtualClusterId(virtualClusterId *string) VirtualClusterInput {
	_init_.Initialize()

	if err := validateVirtualClusterInput_FromVirtualClusterIdParameters(virtualClusterId); err != nil {
		panic(err)
	}
	var returns VirtualClusterInput

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.VirtualClusterInput",
		"fromVirtualClusterId",
		[]interface{}{virtualClusterId},
		&returns,
	)

	return returns
}

