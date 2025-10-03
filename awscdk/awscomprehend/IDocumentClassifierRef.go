package awscomprehend

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscomprehend/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DocumentClassifier.
// Experimental.
type IDocumentClassifierRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDocumentClassifierRef
type jsiiProxy_IDocumentClassifierRef struct {
	internal.Type__constructsIConstruct
}

