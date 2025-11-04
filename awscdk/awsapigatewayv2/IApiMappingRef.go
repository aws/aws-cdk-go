package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiMapping.
// Experimental.
type IApiMappingRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ApiMapping resource.
	// Experimental.
	ApiMappingRef() *ApiMappingReference
}

// The jsii proxy for IApiMappingRef
type jsiiProxy_IApiMappingRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApiMappingRef) ApiMappingRef() *ApiMappingReference {
	var returns *ApiMappingReference
	_jsii_.Get(
		j,
		"apiMappingRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiMappingRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiMappingRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

