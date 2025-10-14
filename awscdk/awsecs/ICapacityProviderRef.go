package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CapacityProvider.
// Experimental.
type ICapacityProviderRef interface {
	constructs.IConstruct
	// A reference to a CapacityProvider resource.
	// Experimental.
	CapacityProviderRef() *CapacityProviderReference
}

// The jsii proxy for ICapacityProviderRef
type jsiiProxy_ICapacityProviderRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICapacityProviderRef) CapacityProviderRef() *CapacityProviderReference {
	var returns *CapacityProviderReference
	_jsii_.Get(
		j,
		"capacityProviderRef",
		&returns,
	)
	return returns
}

