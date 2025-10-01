package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayRouteTable.
// Experimental.
type ILocalGatewayRouteTableRef interface {
	constructs.IConstruct
	// A reference to a LocalGatewayRouteTable resource.
	// Experimental.
	LocalGatewayRouteTableRef() *LocalGatewayRouteTableReference
}

// The jsii proxy for ILocalGatewayRouteTableRef
type jsiiProxy_ILocalGatewayRouteTableRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILocalGatewayRouteTableRef) LocalGatewayRouteTableRef() *LocalGatewayRouteTableReference {
	var returns *LocalGatewayRouteTableReference
	_jsii_.Get(
		j,
		"localGatewayRouteTableRef",
		&returns,
	)
	return returns
}

