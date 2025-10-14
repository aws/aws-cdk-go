package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublicTypeVersion.
// Experimental.
type IPublicTypeVersionRef interface {
	constructs.IConstruct
	// A reference to a PublicTypeVersion resource.
	// Experimental.
	PublicTypeVersionRef() *PublicTypeVersionReference
}

// The jsii proxy for IPublicTypeVersionRef
type jsiiProxy_IPublicTypeVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPublicTypeVersionRef) PublicTypeVersionRef() *PublicTypeVersionReference {
	var returns *PublicTypeVersionReference
	_jsii_.Get(
		j,
		"publicTypeVersionRef",
		&returns,
	)
	return returns
}

