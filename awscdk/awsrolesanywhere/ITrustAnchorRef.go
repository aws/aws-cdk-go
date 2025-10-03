package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrolesanywhere/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustAnchor.
// Experimental.
type ITrustAnchorRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrustAnchorRef
type jsiiProxy_ITrustAnchorRef struct {
	internal.Type__constructsIConstruct
}

