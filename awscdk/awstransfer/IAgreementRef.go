package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Agreement.
// Experimental.
type IAgreementRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAgreementRef
type jsiiProxy_IAgreementRef struct {
	internal.Type__constructsIConstruct
}

