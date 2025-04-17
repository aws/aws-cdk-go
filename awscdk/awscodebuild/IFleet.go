package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
)

// Represents a {@link Fleet} for a reserved capacity CodeBuild project.
type IFleet interface {
	awscdk.IResource
	// The compute type of the fleet.
	// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_codebuild.ComputeType.html
	//
	ComputeType() FleetComputeType
	// The build environment (operating system/architecture/accelerator) type made available to projects using this fleet.
	EnvironmentType() EnvironmentType
	// The ARN of the fleet.
	FleetArn() *string
	// The name of the fleet.
	FleetName() *string
}

// The jsii proxy for IFleet
type jsiiProxy_IFleet struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFleet) ComputeType() FleetComputeType {
	var returns FleetComputeType
	_jsii_.Get(
		j,
		"computeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) EnvironmentType() EnvironmentType {
	var returns EnvironmentType
	_jsii_.Get(
		j,
		"environmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) FleetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) FleetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetName",
		&returns,
	)
	return returns
}

