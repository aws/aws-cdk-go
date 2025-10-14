package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteTable.
// Experimental.
type IRouteTableRef interface {
	constructs.IConstruct
	// A reference to a RouteTable resource.
	// Experimental.
	RouteTableRef() *RouteTableReference
}

// The jsii proxy for IRouteTableRef
type jsiiProxy_IRouteTableRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRouteTableRef) RouteTableRef() *RouteTableReference {
	var returns *RouteTableReference
	_jsii_.Get(
		j,
		"routeTableRef",
		&returns,
	)
	return returns
}

