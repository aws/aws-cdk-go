package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPoint.
// Experimental.
type IAccessPointRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessPointRef
type jsiiProxy_IAccessPointRef struct {
	internal.Type__constructsIConstruct
}

