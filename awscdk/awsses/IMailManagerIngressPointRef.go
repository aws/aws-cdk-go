package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerIngressPoint.
// Experimental.
type IMailManagerIngressPointRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMailManagerIngressPointRef
type jsiiProxy_IMailManagerIngressPointRef struct {
	internal.Type__constructsIConstruct
}

