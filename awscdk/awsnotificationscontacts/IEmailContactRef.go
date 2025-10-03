package awsnotificationscontacts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotificationscontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailContact.
// Experimental.
type IEmailContactRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEmailContactRef
type jsiiProxy_IEmailContactRef struct {
	internal.Type__constructsIConstruct
}

