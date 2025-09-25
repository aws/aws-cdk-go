package awsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EntityType.
// Experimental.
type IEntityTypeRef interface {
	constructs.IConstruct
	// A reference to a EntityType resource.
	// Experimental.
	EntityTypeRef() *EntityTypeReference
}

// The jsii proxy for IEntityTypeRef
type jsiiProxy_IEntityTypeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEntityTypeRef) EntityTypeRef() *EntityTypeReference {
	var returns *EntityTypeReference
	_jsii_.Get(
		j,
		"entityTypeRef",
		&returns,
	)
	return returns
}

