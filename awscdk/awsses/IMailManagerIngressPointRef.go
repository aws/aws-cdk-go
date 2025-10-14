package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerIngressPoint.
// Experimental.
type IMailManagerIngressPointRef interface {
	constructs.IConstruct
	// A reference to a MailManagerIngressPoint resource.
	// Experimental.
	MailManagerIngressPointRef() *MailManagerIngressPointReference
}

// The jsii proxy for IMailManagerIngressPointRef
type jsiiProxy_IMailManagerIngressPointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMailManagerIngressPointRef) MailManagerIngressPointRef() *MailManagerIngressPointReference {
	var returns *MailManagerIngressPointReference
	_jsii_.Get(
		j,
		"mailManagerIngressPointRef",
		&returns,
	)
	return returns
}

