package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ObjectType.
// Experimental.
type IObjectTypeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IObjectTypeRef
type jsiiProxy_IObjectTypeRef struct {
	internal.Type__constructsIConstruct
}

