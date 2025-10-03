package awsopsworks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Instance.
// Experimental.
type IInstanceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInstanceRef
type jsiiProxy_IInstanceRef struct {
	internal.Type__constructsIConstruct
}

