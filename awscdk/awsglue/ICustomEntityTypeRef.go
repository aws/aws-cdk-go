package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomEntityType.
// Experimental.
type ICustomEntityTypeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CustomEntityType resource.
	// Experimental.
	CustomEntityTypeRef() *CustomEntityTypeReference
}

// The jsii proxy for ICustomEntityTypeRef
type jsiiProxy_ICustomEntityTypeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICustomEntityTypeRef) CustomEntityTypeRef() *CustomEntityTypeReference {
	var returns *CustomEntityTypeReference
	_jsii_.Get(
		j,
		"customEntityTypeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomEntityTypeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomEntityTypeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

