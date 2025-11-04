package awsinspectorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspectorv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeSecurityIntegration.
// Experimental.
type ICodeSecurityIntegrationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CodeSecurityIntegration resource.
	// Experimental.
	CodeSecurityIntegrationRef() *CodeSecurityIntegrationReference
}

// The jsii proxy for ICodeSecurityIntegrationRef
type jsiiProxy_ICodeSecurityIntegrationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICodeSecurityIntegrationRef) CodeSecurityIntegrationRef() *CodeSecurityIntegrationReference {
	var returns *CodeSecurityIntegrationReference
	_jsii_.Get(
		j,
		"codeSecurityIntegrationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSecurityIntegrationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSecurityIntegrationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

