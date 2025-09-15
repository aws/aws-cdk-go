package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Location.
// Experimental.
type ILocationRef interface {
	constructs.IConstruct
	// A reference to a Location resource.
	// Experimental.
	LocationRef() *LocationReference
}

// The jsii proxy for ILocationRef
type jsiiProxy_ILocationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILocationRef) LocationRef() *LocationReference {
	var returns *LocationReference
	_jsii_.Get(
		j,
		"locationRef",
		&returns,
	)
	return returns
}

