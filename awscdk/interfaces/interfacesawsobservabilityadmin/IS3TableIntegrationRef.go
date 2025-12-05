package interfacesawsobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a S3TableIntegration.
// Experimental.
type IS3TableIntegrationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a S3TableIntegration resource.
	// Experimental.
	S3TableIntegrationRef() *S3TableIntegrationReference
}

// The jsii proxy for IS3TableIntegrationRef
type jsiiProxy_IS3TableIntegrationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IS3TableIntegrationRef) S3TableIntegrationRef() *S3TableIntegrationReference {
	var returns *S3TableIntegrationReference
	_jsii_.Get(
		j,
		"s3TableIntegrationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IS3TableIntegrationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IS3TableIntegrationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

