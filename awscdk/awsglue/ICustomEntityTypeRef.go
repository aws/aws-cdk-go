package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomEntityType.
// Experimental.
type ICustomEntityTypeRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICustomEntityTypeRef
type jsiiProxy_ICustomEntityTypeRef struct {
	internal.Type__constructsIConstruct
}

