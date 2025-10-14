package awsrolesanywhere

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrolesanywhere/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustAnchor.
// Experimental.
type ITrustAnchorRef interface {
	constructs.IConstruct
	// A reference to a TrustAnchor resource.
	// Experimental.
	TrustAnchorRef() *TrustAnchorReference
}

// The jsii proxy for ITrustAnchorRef
type jsiiProxy_ITrustAnchorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrustAnchorRef) TrustAnchorRef() *TrustAnchorReference {
	var returns *TrustAnchorReference
	_jsii_.Get(
		j,
		"trustAnchorRef",
		&returns,
	)
	return returns
}

