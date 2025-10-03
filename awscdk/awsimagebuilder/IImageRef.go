package awsimagebuilder

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Image.
// Experimental.
type IImageRef interface {
	constructs.IConstruct
}

// The jsii proxy for IImageRef
type jsiiProxy_IImageRef struct {
	internal.Type__constructsIConstruct
}

