package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessGrantsLocation.
// Experimental.
type IAccessGrantsLocationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessGrantsLocationRef
type jsiiProxy_IAccessGrantsLocationRef struct {
	internal.Type__constructsIConstruct
}

