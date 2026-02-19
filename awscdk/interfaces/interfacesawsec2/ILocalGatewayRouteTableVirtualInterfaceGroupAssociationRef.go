package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayRouteTableVirtualInterfaceGroupAssociation.
// Experimental.
type ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocalGatewayRouteTableVirtualInterfaceGroupAssociation resource.
	// Experimental.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociationRef() *LocalGatewayRouteTableVirtualInterfaceGroupAssociationReference
}

// The jsii proxy for ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef
type jsiiProxy_ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef) LocalGatewayRouteTableVirtualInterfaceGroupAssociationRef() *LocalGatewayRouteTableVirtualInterfaceGroupAssociationReference {
	var returns *LocalGatewayRouteTableVirtualInterfaceGroupAssociationReference
	_jsii_.Get(
		j,
		"localGatewayRouteTableVirtualInterfaceGroupAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

