package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerTrafficPolicy.
// Experimental.
type IMailManagerTrafficPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMailManagerTrafficPolicyRef
type jsiiProxy_IMailManagerTrafficPolicyRef struct {
	internal.Type__constructsIConstruct
}

