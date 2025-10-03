package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Classifier.
// Experimental.
type IClassifierRef interface {
	constructs.IConstruct
}

// The jsii proxy for IClassifierRef
type jsiiProxy_IClassifierRef struct {
	internal.Type__constructsIConstruct
}

