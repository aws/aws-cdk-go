package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcOrigin.
// Experimental.
type IVpcOriginRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VpcOrigin resource.
	// Experimental.
	VpcOriginRef() *VpcOriginReference
}

// The jsii proxy for IVpcOriginRef
type jsiiProxy_IVpcOriginRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVpcOriginRef) VpcOriginRef() *VpcOriginReference {
	var returns *VpcOriginReference
	_jsii_.Get(
		j,
		"vpcOriginRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcOriginRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcOriginRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

