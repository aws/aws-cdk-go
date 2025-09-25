package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerRuleSet.
// Experimental.
type IMailManagerRuleSetRef interface {
	constructs.IConstruct
	// A reference to a MailManagerRuleSet resource.
	// Experimental.
	MailManagerRuleSetRef() *MailManagerRuleSetReference
}

// The jsii proxy for IMailManagerRuleSetRef
type jsiiProxy_IMailManagerRuleSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMailManagerRuleSetRef) MailManagerRuleSetRef() *MailManagerRuleSetReference {
	var returns *MailManagerRuleSetReference
	_jsii_.Get(
		j,
		"mailManagerRuleSetRef",
		&returns,
	)
	return returns
}

