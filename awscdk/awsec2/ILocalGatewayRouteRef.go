package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayRoute.
// Experimental.
type ILocalGatewayRouteRef interface {
	constructs.IConstruct
	// A reference to a LocalGatewayRoute resource.
	// Experimental.
	LocalGatewayRouteRef() *LocalGatewayRouteReference
}

// The jsii proxy for ILocalGatewayRouteRef
type jsiiProxy_ILocalGatewayRouteRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILocalGatewayRouteRef) LocalGatewayRouteRef() *LocalGatewayRouteReference {
	var returns *LocalGatewayRouteReference
	_jsii_.Get(
		j,
		"localGatewayRouteRef",
		&returns,
	)
	return returns
}

