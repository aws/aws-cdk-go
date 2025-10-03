package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DNSSEC.
// Experimental.
type IDNSSECRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDNSSECRef
type jsiiProxy_IDNSSECRef struct {
	internal.Type__constructsIConstruct
}

