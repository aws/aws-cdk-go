package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerRuleSet.
// Experimental.
type IMailManagerRuleSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMailManagerRuleSetRef
type jsiiProxy_IMailManagerRuleSetRef struct {
	internal.Type__constructsIConstruct
}

