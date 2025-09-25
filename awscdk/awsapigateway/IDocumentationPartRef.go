package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DocumentationPart.
// Experimental.
type IDocumentationPartRef interface {
	constructs.IConstruct
	// A reference to a DocumentationPart resource.
	// Experimental.
	DocumentationPartRef() *DocumentationPartReference
}

// The jsii proxy for IDocumentationPartRef
type jsiiProxy_IDocumentationPartRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDocumentationPartRef) DocumentationPartRef() *DocumentationPartReference {
	var returns *DocumentationPartReference
	_jsii_.Get(
		j,
		"documentationPartRef",
		&returns,
	)
	return returns
}

