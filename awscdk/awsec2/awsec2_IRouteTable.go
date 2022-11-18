package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An abstract route table.
type IRouteTable interface {
	// Route table ID.
	RouteTableId() *string
}

// The jsii proxy for IRouteTable
type jsiiProxy_IRouteTable struct {
	_ byte // padding
}

func (j *jsiiProxy_IRouteTable) RouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableId",
		&returns,
	)
	return returns
}

