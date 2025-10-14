package awsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OdbNetwork.
// Experimental.
type IOdbNetworkRef interface {
	constructs.IConstruct
	// A reference to a OdbNetwork resource.
	// Experimental.
	OdbNetworkRef() *OdbNetworkReference
}

// The jsii proxy for IOdbNetworkRef
type jsiiProxy_IOdbNetworkRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOdbNetworkRef) OdbNetworkRef() *OdbNetworkReference {
	var returns *OdbNetworkReference
	_jsii_.Get(
		j,
		"odbNetworkRef",
		&returns,
	)
	return returns
}

