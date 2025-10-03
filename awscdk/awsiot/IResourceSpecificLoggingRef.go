package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceSpecificLogging.
// Experimental.
type IResourceSpecificLoggingRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceSpecificLoggingRef
type jsiiProxy_IResourceSpecificLoggingRef struct {
	internal.Type__constructsIConstruct
}

