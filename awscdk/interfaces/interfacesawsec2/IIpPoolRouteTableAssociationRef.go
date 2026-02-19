package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IpPoolRouteTableAssociation.
// Experimental.
type IIpPoolRouteTableAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IpPoolRouteTableAssociation resource.
	// Experimental.
	IpPoolRouteTableAssociationRef() *IpPoolRouteTableAssociationReference
}

// The jsii proxy for IIpPoolRouteTableAssociationRef
type jsiiProxy_IIpPoolRouteTableAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIpPoolRouteTableAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIpPoolRouteTableAssociationRef) IpPoolRouteTableAssociationRef() *IpPoolRouteTableAssociationReference {
	var returns *IpPoolRouteTableAssociationReference
	_jsii_.Get(
		j,
		"ipPoolRouteTableAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpPoolRouteTableAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpPoolRouteTableAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

