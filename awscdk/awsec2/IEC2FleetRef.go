package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EC2Fleet.
// Experimental.
type IEC2FleetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EC2Fleet resource.
	// Experimental.
	Ec2FleetRef() *EC2FleetReference
}

// The jsii proxy for IEC2FleetRef
type jsiiProxy_IEC2FleetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEC2FleetRef) Ec2FleetRef() *EC2FleetReference {
	var returns *EC2FleetReference
	_jsii_.Get(
		j,
		"ec2FleetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEC2FleetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEC2FleetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

