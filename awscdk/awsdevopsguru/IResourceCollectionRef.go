package awsdevopsguru

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevopsguru/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceCollection.
// Experimental.
type IResourceCollectionRef interface {
	constructs.IConstruct
	// A reference to a ResourceCollection resource.
	// Experimental.
	ResourceCollectionRef() *ResourceCollectionReference
}

// The jsii proxy for IResourceCollectionRef
type jsiiProxy_IResourceCollectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceCollectionRef) ResourceCollectionRef() *ResourceCollectionReference {
	var returns *ResourceCollectionReference
	_jsii_.Get(
		j,
		"resourceCollectionRef",
		&returns,
	)
	return returns
}

