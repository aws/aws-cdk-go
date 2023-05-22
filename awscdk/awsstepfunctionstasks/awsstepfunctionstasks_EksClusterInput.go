package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awseks"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Class that supports methods which return the EKS cluster name depending on input type.
//
// Example:
//   tasks.NewEmrContainersCreateVirtualCluster(this, jsii.String("Create a Virtual Cluster"), &EmrContainersCreateVirtualClusterProps{
//   	EksCluster: tasks.EksClusterInput_FromTaskInput(sfn.TaskInput_FromText(jsii.String("clusterId"))),
//   	EksNamespace: jsii.String("specified-namespace"),
//   })
//
// Experimental.
type EksClusterInput interface {
	// The name of the EKS Cluster.
	// Experimental.
	ClusterName() *string
}

// The jsii proxy struct for EksClusterInput
type jsiiProxy_EksClusterInput struct {
	_ byte // padding
}

func (j *jsiiProxy_EksClusterInput) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}


// Specify an existing EKS Cluster as the name for this Cluster.
// Experimental.
func EksClusterInput_FromCluster(cluster awseks.ICluster) EksClusterInput {
	_init_.Initialize()

	if err := validateEksClusterInput_FromClusterParameters(cluster); err != nil {
		panic(err)
	}
	var returns EksClusterInput

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EksClusterInput",
		"fromCluster",
		[]interface{}{cluster},
		&returns,
	)

	return returns
}

// Specify a Task Input as the name for this Cluster.
// Experimental.
func EksClusterInput_FromTaskInput(taskInput awsstepfunctions.TaskInput) EksClusterInput {
	_init_.Initialize()

	if err := validateEksClusterInput_FromTaskInputParameters(taskInput); err != nil {
		panic(err)
	}
	var returns EksClusterInput

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EksClusterInput",
		"fromTaskInput",
		[]interface{}{taskInput},
		&returns,
	)

	return returns
}

