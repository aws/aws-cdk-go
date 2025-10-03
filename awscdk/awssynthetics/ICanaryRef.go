package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Canary.
// Experimental.
type ICanaryRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICanaryRef
type jsiiProxy_ICanaryRef struct {
	internal.Type__constructsIConstruct
}

