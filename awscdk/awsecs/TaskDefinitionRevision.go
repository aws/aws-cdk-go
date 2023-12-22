package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents revision of a task definition, either a specific numbered revision or the `latest` revision.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//
//
//   ecs.NewExternalService(this, jsii.String("Service"), &ExternalServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	DesiredCount: jsii.Number(5),
//   	TaskDefinitionRevision: ecs.TaskDefinitionRevision_Of(jsii.Number(1)),
//   })
//
//   ecs.NewExternalService(this, jsii.String("Service"), &ExternalServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	DesiredCount: jsii.Number(5),
//   	TaskDefinitionRevision: ecs.TaskDefinitionRevision_LATEST(),
//   })
//
type TaskDefinitionRevision interface {
	// The string representation of this revision.
	Revision() *string
}

// The jsii proxy struct for TaskDefinitionRevision
type jsiiProxy_TaskDefinitionRevision struct {
	_ byte // padding
}

func (j *jsiiProxy_TaskDefinitionRevision) Revision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}


// Specific revision of a task.
func TaskDefinitionRevision_Of(revision *float64) TaskDefinitionRevision {
	_init_.Initialize()

	if err := validateTaskDefinitionRevision_OfParameters(revision); err != nil {
		panic(err)
	}
	var returns TaskDefinitionRevision

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinitionRevision",
		"of",
		[]interface{}{revision},
		&returns,
	)

	return returns
}

func TaskDefinitionRevision_LATEST() TaskDefinitionRevision {
	_init_.Initialize()
	var returns TaskDefinitionRevision
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.TaskDefinitionRevision",
		"LATEST",
		&returns,
	)
	return returns
}

