package awsresourceexplorer2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsresourceexplorer2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a View.
// Experimental.
type IViewRef interface {
	constructs.IConstruct
}

// The jsii proxy for IViewRef
type jsiiProxy_IViewRef struct {
	internal.Type__constructsIConstruct
}

