package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcIngressConnection.
// Experimental.
type IVpcIngressConnectionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VpcIngressConnection resource.
	// Experimental.
	VpcIngressConnectionRef() *VpcIngressConnectionReference
}

// The jsii proxy for IVpcIngressConnectionRef
type jsiiProxy_IVpcIngressConnectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVpcIngressConnectionRef) VpcIngressConnectionRef() *VpcIngressConnectionReference {
	var returns *VpcIngressConnectionReference
	_jsii_.Get(
		j,
		"vpcIngressConnectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcIngressConnectionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcIngressConnectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

