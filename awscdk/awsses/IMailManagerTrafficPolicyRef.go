package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerTrafficPolicy.
// Experimental.
type IMailManagerTrafficPolicyRef interface {
	constructs.IConstruct
	// A reference to a MailManagerTrafficPolicy resource.
	// Experimental.
	MailManagerTrafficPolicyRef() *MailManagerTrafficPolicyReference
}

// The jsii proxy for IMailManagerTrafficPolicyRef
type jsiiProxy_IMailManagerTrafficPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMailManagerTrafficPolicyRef) MailManagerTrafficPolicyRef() *MailManagerTrafficPolicyReference {
	var returns *MailManagerTrafficPolicyReference
	_jsii_.Get(
		j,
		"mailManagerTrafficPolicyRef",
		&returns,
	)
	return returns
}

