package awslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APIKey.
// Experimental.
type IAPIKeyRef interface {
	constructs.IConstruct
	// A reference to a APIKey resource.
	// Experimental.
	ApiKeyRef() *APIKeyReference
}

// The jsii proxy for IAPIKeyRef
type jsiiProxy_IAPIKeyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAPIKeyRef) ApiKeyRef() *APIKeyReference {
	var returns *APIKeyReference
	_jsii_.Get(
		j,
		"apiKeyRef",
		&returns,
	)
	return returns
}

