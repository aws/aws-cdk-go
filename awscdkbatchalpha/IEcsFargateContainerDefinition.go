package awscdkbatchalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// A container orchestrated by ECS that uses Fargate resources and is orchestrated by ECS.
// Experimental.
type IEcsFargateContainerDefinition interface {
	IEcsContainerDefinition
	// Indicates whether the job has a public IP address.
	//
	// For a job that's running on Fargate resources in a private subnet to send outbound traffic to the internet
	// (for example, to pull container images), the private subnet requires a NAT gateway be attached to route requests to the internet.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html
	//
	// Experimental.
	AssignPublicIp() *bool
	// Which version of Fargate to use when running this container.
	// Experimental.
	FargatePlatformVersion() awsecs.FargatePlatformVersion
}

// The jsii proxy for IEcsFargateContainerDefinition
type jsiiProxy_IEcsFargateContainerDefinition struct {
	jsiiProxy_IEcsContainerDefinition
}

func (j *jsiiProxy_IEcsFargateContainerDefinition) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsFargateContainerDefinition) FargatePlatformVersion() awsecs.FargatePlatformVersion {
	var returns awsecs.FargatePlatformVersion
	_jsii_.Get(
		j,
		"fargatePlatformVersion",
		&returns,
	)
	return returns
}

