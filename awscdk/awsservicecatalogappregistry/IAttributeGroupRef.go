package awsservicecatalogappregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalogappregistry/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AttributeGroup.
// Experimental.
type IAttributeGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AttributeGroup resource.
	// Experimental.
	AttributeGroupRef() *AttributeGroupReference
}

// The jsii proxy for IAttributeGroupRef
type jsiiProxy_IAttributeGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAttributeGroupRef) AttributeGroupRef() *AttributeGroupReference {
	var returns *AttributeGroupReference
	_jsii_.Get(
		j,
		"attributeGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAttributeGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAttributeGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

