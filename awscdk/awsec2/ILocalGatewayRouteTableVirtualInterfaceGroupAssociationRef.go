package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayRouteTableVirtualInterfaceGroupAssociation.
// Experimental.
type ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef interface {
	constructs.IConstruct
	// A reference to a LocalGatewayRouteTableVirtualInterfaceGroupAssociation resource.
	// Experimental.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociationRef() *LocalGatewayRouteTableVirtualInterfaceGroupAssociationReference
}

// The jsii proxy for ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef
type jsiiProxy_ILocalGatewayRouteTableVirtualInterfaceGroupAssociationRef struct {
	internal.Type__constructsIConstruct
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

