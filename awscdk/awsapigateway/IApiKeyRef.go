package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiKey.
// Experimental.
type IApiKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApiKeyRef
type jsiiProxy_IApiKeyRef struct {
	internal.Type__constructsIConstruct
}

