package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegistryPolicy.
// Experimental.
type IRegistryPolicyRef interface {
	constructs.IConstruct
	// A reference to a RegistryPolicy resource.
	// Experimental.
	RegistryPolicyRef() *RegistryPolicyReference
}

// The jsii proxy for IRegistryPolicyRef
type jsiiProxy_IRegistryPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRegistryPolicyRef) RegistryPolicyRef() *RegistryPolicyReference {
	var returns *RegistryPolicyReference
	_jsii_.Get(
		j,
		"registryPolicyRef",
		&returns,
	)
	return returns
}

