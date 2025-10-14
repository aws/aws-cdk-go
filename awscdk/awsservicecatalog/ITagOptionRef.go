package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagOption.
// Experimental.
type ITagOptionRef interface {
	constructs.IConstruct
	// A reference to a TagOption resource.
	// Experimental.
	TagOptionRef() *TagOptionReference
}

// The jsii proxy for ITagOptionRef
type jsiiProxy_ITagOptionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITagOptionRef) TagOptionRef() *TagOptionReference {
	var returns *TagOptionReference
	_jsii_.Get(
		j,
		"tagOptionRef",
		&returns,
	)
	return returns
}

