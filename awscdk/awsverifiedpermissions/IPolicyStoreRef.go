package awsverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsverifiedpermissions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyStore.
// Experimental.
type IPolicyStoreRef interface {
	constructs.IConstruct
	// A reference to a PolicyStore resource.
	// Experimental.
	PolicyStoreRef() *PolicyStoreReference
}

// The jsii proxy for IPolicyStoreRef
type jsiiProxy_IPolicyStoreRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPolicyStoreRef) PolicyStoreRef() *PolicyStoreReference {
	var returns *PolicyStoreReference
	_jsii_.Get(
		j,
		"policyStoreRef",
		&returns,
	)
	return returns
}

