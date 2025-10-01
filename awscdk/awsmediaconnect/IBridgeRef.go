package awsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Bridge.
// Experimental.
type IBridgeRef interface {
	constructs.IConstruct
	// A reference to a Bridge resource.
	// Experimental.
	BridgeRef() *BridgeReference
}

// The jsii proxy for IBridgeRef
type jsiiProxy_IBridgeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBridgeRef) BridgeRef() *BridgeReference {
	var returns *BridgeReference
	_jsii_.Get(
		j,
		"bridgeRef",
		&returns,
	)
	return returns
}

