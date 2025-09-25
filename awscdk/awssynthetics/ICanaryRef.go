package awssynthetics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Canary.
// Experimental.
type ICanaryRef interface {
	constructs.IConstruct
	// A reference to a Canary resource.
	// Experimental.
	CanaryRef() *CanaryReference
}

// The jsii proxy for ICanaryRef
type jsiiProxy_ICanaryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICanaryRef) CanaryRef() *CanaryReference {
	var returns *CanaryReference
	_jsii_.Get(
		j,
		"canaryRef",
		&returns,
	)
	return returns
}

