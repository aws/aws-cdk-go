package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThreatIntelSet.
// Experimental.
type IThreatIntelSetRef interface {
	constructs.IConstruct
	// A reference to a ThreatIntelSet resource.
	// Experimental.
	ThreatIntelSetRef() *ThreatIntelSetReference
}

// The jsii proxy for IThreatIntelSetRef
type jsiiProxy_IThreatIntelSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IThreatIntelSetRef) ThreatIntelSetRef() *ThreatIntelSetReference {
	var returns *ThreatIntelSetReference
	_jsii_.Get(
		j,
		"threatIntelSetRef",
		&returns,
	)
	return returns
}

