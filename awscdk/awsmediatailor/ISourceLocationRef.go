package awsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SourceLocation.
// Experimental.
type ISourceLocationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SourceLocation resource.
	// Experimental.
	SourceLocationRef() *SourceLocationReference
}

// The jsii proxy for ISourceLocationRef
type jsiiProxy_ISourceLocationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISourceLocationRef) SourceLocationRef() *SourceLocationReference {
	var returns *SourceLocationReference
	_jsii_.Get(
		j,
		"sourceLocationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceLocationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceLocationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

