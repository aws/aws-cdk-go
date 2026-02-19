package interfacesawscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcOrigin.
// Experimental.
type IVpcOriginRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VpcOrigin resource.
	// Experimental.
	VpcOriginRef() *VpcOriginReference
}

// The jsii proxy for IVpcOriginRef
type jsiiProxy_IVpcOriginRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVpcOriginRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IVpcOriginRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

