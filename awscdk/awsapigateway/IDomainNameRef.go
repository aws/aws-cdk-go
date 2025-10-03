package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainName.
// Experimental.
type IDomainNameRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDomainNameRef
type jsiiProxy_IDomainNameRef struct {
	internal.Type__constructsIConstruct
}

