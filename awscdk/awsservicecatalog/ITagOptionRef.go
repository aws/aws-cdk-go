package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagOption.
// Experimental.
type ITagOptionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITagOptionRef
type jsiiProxy_ITagOptionRef struct {
	internal.Type__constructsIConstruct
}

