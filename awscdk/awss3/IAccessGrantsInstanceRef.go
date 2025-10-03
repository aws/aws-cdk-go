package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessGrantsInstance.
// Experimental.
type IAccessGrantsInstanceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessGrantsInstanceRef
type jsiiProxy_IAccessGrantsInstanceRef struct {
	internal.Type__constructsIConstruct
}

