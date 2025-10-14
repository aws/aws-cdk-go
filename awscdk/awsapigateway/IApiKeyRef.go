package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiKey.
// Experimental.
type IApiKeyRef interface {
	constructs.IConstruct
	// A reference to a ApiKey resource.
	// Experimental.
	ApiKeyRef() *ApiKeyReference
}

// The jsii proxy for IApiKeyRef
type jsiiProxy_IApiKeyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApiKeyRef) ApiKeyRef() *ApiKeyReference {
	var returns *ApiKeyReference
	_jsii_.Get(
		j,
		"apiKeyRef",
		&returns,
	)
	return returns
}

