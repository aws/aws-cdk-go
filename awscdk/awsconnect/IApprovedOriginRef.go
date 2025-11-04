package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApprovedOrigin.
// Experimental.
type IApprovedOriginRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ApprovedOrigin resource.
	// Experimental.
	ApprovedOriginRef() *ApprovedOriginReference
}

// The jsii proxy for IApprovedOriginRef
type jsiiProxy_IApprovedOriginRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApprovedOriginRef) ApprovedOriginRef() *ApprovedOriginReference {
	var returns *ApprovedOriginReference
	_jsii_.Get(
		j,
		"approvedOriginRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApprovedOriginRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApprovedOriginRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

