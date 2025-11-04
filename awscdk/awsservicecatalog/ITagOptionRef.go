package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagOption.
// Experimental.
type ITagOptionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TagOption resource.
	// Experimental.
	TagOptionRef() *TagOptionReference
}

// The jsii proxy for ITagOptionRef
type jsiiProxy_ITagOptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ITagOptionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagOptionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

