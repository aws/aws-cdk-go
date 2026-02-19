package interfacesawsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcEndpoint.
// Experimental.
type IVpcEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VpcEndpoint resource.
	// Experimental.
	VpcEndpointRef() *VpcEndpointReference
}

// The jsii proxy for IVpcEndpointRef
type jsiiProxy_IVpcEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVpcEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVpcEndpointRef) VpcEndpointRef() *VpcEndpointReference {
	var returns *VpcEndpointReference
	_jsii_.Get(
		j,
		"vpcEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

