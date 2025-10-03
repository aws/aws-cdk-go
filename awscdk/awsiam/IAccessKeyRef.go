package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessKey.
// Experimental.
type IAccessKeyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessKeyRef
type jsiiProxy_IAccessKeyRef struct {
	internal.Type__constructsIConstruct
}

