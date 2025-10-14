package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Secret.
// Experimental.
type ISecretRef interface {
	constructs.IConstruct
	// A reference to a Secret resource.
	// Experimental.
	SecretRef() *SecretReference
}

// The jsii proxy for ISecretRef
type jsiiProxy_ISecretRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecretRef) SecretRef() *SecretReference {
	var returns *SecretReference
	_jsii_.Get(
		j,
		"secretRef",
		&returns,
	)
	return returns
}

