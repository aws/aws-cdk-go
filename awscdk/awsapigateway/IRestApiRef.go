package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RestApi.
// Experimental.
type IRestApiRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRestApiRef
type jsiiProxy_IRestApiRef struct {
	internal.Type__constructsIConstruct
}

