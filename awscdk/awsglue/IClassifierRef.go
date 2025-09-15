package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Classifier.
// Experimental.
type IClassifierRef interface {
	constructs.IConstruct
	// A reference to a Classifier resource.
	// Experimental.
	ClassifierRef() *ClassifierReference
}

// The jsii proxy for IClassifierRef
type jsiiProxy_IClassifierRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IClassifierRef) ClassifierRef() *ClassifierReference {
	var returns *ClassifierReference
	_jsii_.Get(
		j,
		"classifierRef",
		&returns,
	)
	return returns
}

