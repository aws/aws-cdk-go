package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomDataIdentifier.
// Experimental.
type ICustomDataIdentifierRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CustomDataIdentifier resource.
	// Experimental.
	CustomDataIdentifierRef() *CustomDataIdentifierReference
}

// The jsii proxy for ICustomDataIdentifierRef
type jsiiProxy_ICustomDataIdentifierRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICustomDataIdentifierRef) CustomDataIdentifierRef() *CustomDataIdentifierReference {
	var returns *CustomDataIdentifierReference
	_jsii_.Get(
		j,
		"customDataIdentifierRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomDataIdentifierRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomDataIdentifierRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

