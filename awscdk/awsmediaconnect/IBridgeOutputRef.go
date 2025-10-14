package awsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BridgeOutput.
// Experimental.
type IBridgeOutputRef interface {
	constructs.IConstruct
	// A reference to a BridgeOutput resource.
	// Experimental.
	BridgeOutputRef() *BridgeOutputReference
}

// The jsii proxy for IBridgeOutputRef
type jsiiProxy_IBridgeOutputRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBridgeOutputRef) BridgeOutputRef() *BridgeOutputReference {
	var returns *BridgeOutputReference
	_jsii_.Get(
		j,
		"bridgeOutputRef",
		&returns,
	)
	return returns
}

