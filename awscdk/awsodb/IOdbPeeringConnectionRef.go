package awsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OdbPeeringConnection.
// Experimental.
type IOdbPeeringConnectionRef interface {
	constructs.IConstruct
	// A reference to a OdbPeeringConnection resource.
	// Experimental.
	OdbPeeringConnectionRef() *OdbPeeringConnectionReference
}

// The jsii proxy for IOdbPeeringConnectionRef
type jsiiProxy_IOdbPeeringConnectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOdbPeeringConnectionRef) OdbPeeringConnectionRef() *OdbPeeringConnectionReference {
	var returns *OdbPeeringConnectionReference
	_jsii_.Get(
		j,
		"odbPeeringConnectionRef",
		&returns,
	)
	return returns
}

