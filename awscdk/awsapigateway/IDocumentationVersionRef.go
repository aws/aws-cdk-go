package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DocumentationVersion.
// Experimental.
type IDocumentationVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDocumentationVersionRef
type jsiiProxy_IDocumentationVersionRef struct {
	internal.Type__constructsIConstruct
}

