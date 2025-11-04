package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Api.
// Experimental.
type IApiRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Api resource.
	// Experimental.
	ApiRef() *ApiReference
}

// The jsii proxy for IApiRef
type jsiiProxy_IApiRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApiRef) ApiRef() *ApiReference {
	var returns *ApiReference
	_jsii_.Get(
		j,
		"apiRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

