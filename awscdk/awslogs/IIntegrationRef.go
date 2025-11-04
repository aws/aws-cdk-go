package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Integration.
// Experimental.
type IIntegrationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Integration resource.
	// Experimental.
	IntegrationRef() *IntegrationReference
}

// The jsii proxy for IIntegrationRef
type jsiiProxy_IIntegrationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IIntegrationRef) IntegrationRef() *IntegrationReference {
	var returns *IntegrationReference
	_jsii_.Get(
		j,
		"integrationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntegrationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntegrationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

