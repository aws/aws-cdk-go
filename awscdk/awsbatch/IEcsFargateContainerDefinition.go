package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// A container orchestrated by ECS that uses Fargate resources and is orchestrated by ECS.
type IEcsFargateContainerDefinition interface {
	IEcsContainerDefinition
	// Indicates whether the job has a public IP address.
	//
	// For a job that's running on Fargate resources in a private subnet to send outbound traffic to the internet
	// (for example, to pull container images), the private subnet requires a NAT gateway be attached to route requests to the internet.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html
	//
	// Default: false.
	//
	AssignPublicIp() *bool
	// The size for ephemeral storage.
	// Default: - 20 GiB.
	//
	EphemeralStorageSize() awscdk.Size
	// The vCPU architecture of Fargate Runtime.
	// Default: - X86_64.
	//
	FargateCpuArchitecture() awsecs.CpuArchitecture
	// The operating system for the compute environment.
	// Default: - LINUX.
	//
	FargateOperatingSystemFamily() awsecs.OperatingSystemFamily
	// Which version of Fargate to use when running this container.
	// Default: LATEST.
	//
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

func (j *jsiiProxy_IEcsFargateContainerDefinition) EphemeralStorageSize() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"ephemeralStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsFargateContainerDefinition) FargateCpuArchitecture() awsecs.CpuArchitecture {
	var returns awsecs.CpuArchitecture
	_jsii_.Get(
		j,
		"fargateCpuArchitecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsFargateContainerDefinition) FargateOperatingSystemFamily() awsecs.OperatingSystemFamily {
	var returns awsecs.OperatingSystemFamily
	_jsii_.Get(
		j,
		"fargateOperatingSystemFamily",
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

