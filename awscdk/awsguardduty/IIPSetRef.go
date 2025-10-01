package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPSet.
// Experimental.
type IIPSetRef interface {
	constructs.IConstruct
	// A reference to a IPSet resource.
	// Experimental.
	IpSetRef() *IPSetReference
}

// The jsii proxy for IIPSetRef
type jsiiProxy_IIPSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIPSetRef) IpSetRef() *IPSetReference {
	var returns *IPSetReference
	_jsii_.Get(
		j,
		"ipSetRef",
		&returns,
	)
	return returns
}

