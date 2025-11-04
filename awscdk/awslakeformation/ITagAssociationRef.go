package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagAssociation.
// Experimental.
type ITagAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TagAssociation resource.
	// Experimental.
	TagAssociationRef() *TagAssociationReference
}

// The jsii proxy for ITagAssociationRef
type jsiiProxy_ITagAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITagAssociationRef) TagAssociationRef() *TagAssociationReference {
	var returns *TagAssociationReference
	_jsii_.Get(
		j,
		"tagAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

