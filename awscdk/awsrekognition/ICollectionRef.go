package awsrekognition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrekognition/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Collection.
// Experimental.
type ICollectionRef interface {
	constructs.IConstruct
	// A reference to a Collection resource.
	// Experimental.
	CollectionRef() *CollectionReference
}

// The jsii proxy for ICollectionRef
type jsiiProxy_ICollectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICollectionRef) CollectionRef() *CollectionReference {
	var returns *CollectionReference
	_jsii_.Get(
		j,
		"collectionRef",
		&returns,
	)
	return returns
}

