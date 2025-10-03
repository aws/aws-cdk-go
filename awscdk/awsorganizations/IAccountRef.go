package awsorganizations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsorganizations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Account.
// Experimental.
type IAccountRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccountRef
type jsiiProxy_IAccountRef struct {
	internal.Type__constructsIConstruct
}

