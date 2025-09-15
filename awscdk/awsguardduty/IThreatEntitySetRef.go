package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThreatEntitySet.
// Experimental.
type IThreatEntitySetRef interface {
	constructs.IConstruct
	// A reference to a ThreatEntitySet resource.
	// Experimental.
	ThreatEntitySetRef() *ThreatEntitySetReference
}

// The jsii proxy for IThreatEntitySetRef
type jsiiProxy_IThreatEntitySetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IThreatEntitySetRef) ThreatEntitySetRef() *ThreatEntitySetReference {
	var returns *ThreatEntitySetReference
	_jsii_.Get(
		j,
		"threatEntitySetRef",
		&returns,
	)
	return returns
}

