package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Account.
// Experimental.
type IAccountRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccountRef
type jsiiProxy_IAccountRef struct {
	internal.Type__constructsIConstruct
}

