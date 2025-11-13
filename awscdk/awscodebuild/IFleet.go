package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Fleet for a reserved capacity CodeBuild project.
type IFleet interface {
	awsec2.IConnectable
	awsiam.IGrantable
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
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFleet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IFleet) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

