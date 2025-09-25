package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DNSSEC.
// Experimental.
type IDNSSECRef interface {
	constructs.IConstruct
	// A reference to a DNSSEC resource.
	// Experimental.
	DnssecRef() *DNSSECReference
}

// The jsii proxy for IDNSSECRef
type jsiiProxy_IDNSSECRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDNSSECRef) DnssecRef() *DNSSECReference {
	var returns *DNSSECReference
	_jsii_.Get(
		j,
		"dnssecRef",
		&returns,
	)
	return returns
}

