package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MLTransform.
// Experimental.
type IMLTransformRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMLTransformRef
type jsiiProxy_IMLTransformRef struct {
	internal.Type__constructsIConstruct
}

