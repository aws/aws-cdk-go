package awstransfer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Agreement.
// Experimental.
type IAgreementRef interface {
	constructs.IConstruct
	// A reference to a Agreement resource.
	// Experimental.
	AgreementRef() *AgreementReference
}

// The jsii proxy for IAgreementRef
type jsiiProxy_IAgreementRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAgreementRef) AgreementRef() *AgreementReference {
	var returns *AgreementReference
	_jsii_.Get(
		j,
		"agreementRef",
		&returns,
	)
	return returns
}

