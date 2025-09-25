package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Publisher.
// Experimental.
type IPublisherRef interface {
	constructs.IConstruct
	// A reference to a Publisher resource.
	// Experimental.
	PublisherRef() *PublisherReference
}

// The jsii proxy for IPublisherRef
type jsiiProxy_IPublisherRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPublisherRef) PublisherRef() *PublisherReference {
	var returns *PublisherReference
	_jsii_.Get(
		j,
		"publisherRef",
		&returns,
	)
	return returns
}

