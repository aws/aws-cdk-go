package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainNameV2.
// Experimental.
type IDomainNameV2Ref interface {
	constructs.IConstruct
}

// The jsii proxy for IDomainNameV2Ref
type jsiiProxy_IDomainNameV2Ref struct {
	internal.Type__constructsIConstruct
}

