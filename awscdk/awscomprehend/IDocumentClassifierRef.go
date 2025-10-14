package awscomprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscomprehend/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DocumentClassifier.
// Experimental.
type IDocumentClassifierRef interface {
	constructs.IConstruct
	// A reference to a DocumentClassifier resource.
	// Experimental.
	DocumentClassifierRef() *DocumentClassifierReference
}

// The jsii proxy for IDocumentClassifierRef
type jsiiProxy_IDocumentClassifierRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDocumentClassifierRef) DocumentClassifierRef() *DocumentClassifierReference {
	var returns *DocumentClassifierReference
	_jsii_.Get(
		j,
		"documentClassifierRef",
		&returns,
	)
	return returns
}

