package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomDBEngineVersion.
// Experimental.
type ICustomDBEngineVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CustomDBEngineVersion resource.
	// Experimental.
	CustomDbEngineVersionRef() *CustomDBEngineVersionReference
}

// The jsii proxy for ICustomDBEngineVersionRef
type jsiiProxy_ICustomDBEngineVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICustomDBEngineVersionRef) CustomDbEngineVersionRef() *CustomDBEngineVersionReference {
	var returns *CustomDBEngineVersionReference
	_jsii_.Get(
		j,
		"customDbEngineVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomDBEngineVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomDBEngineVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

