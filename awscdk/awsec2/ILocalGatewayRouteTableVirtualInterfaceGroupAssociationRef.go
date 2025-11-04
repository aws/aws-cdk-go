package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayRouteTableVirtualInterfaceGroupAssociation.
// Experimental.
type ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LocalGatewayRouteTableVirtualInterfaceGroupAssociation resource.
	// Experimental.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociationRef() *LocalGatewayRouteTableVirtualInterfaceGroupAssociationReference
}

// The jsii proxy for ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef
type jsiiProxy_ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

