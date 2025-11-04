package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayRouteTableVPCAssociation.
// Experimental.
type ILocalGatewayRouteTableVPCAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LocalGatewayRouteTableVPCAssociation resource.
	// Experimental.
	LocalGatewayRouteTableVpcAssociationRef() *LocalGatewayRouteTableVPCAssociationReference
}

// The jsii proxy for ILocalGatewayRouteTableVPCAssociationRef
type jsiiProxy_ILocalGatewayRouteTableVPCAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILocalGatewayRouteTableVPCAssociationRef) LocalGatewayRouteTableVpcAssociationRef() *LocalGatewayRouteTableVPCAssociationReference {
	var returns *LocalGatewayRouteTableVPCAssociationReference
	_jsii_.Get(
		j,
		"localGatewayRouteTableVpcAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayRouteTableVPCAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayRouteTableVPCAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

