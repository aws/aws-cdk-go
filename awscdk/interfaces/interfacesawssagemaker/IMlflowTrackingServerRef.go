package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MlflowTrackingServer.
// Experimental.
type IMlflowTrackingServerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MlflowTrackingServer resource.
	// Experimental.
	MlflowTrackingServerRef() *MlflowTrackingServerReference
}

// The jsii proxy for IMlflowTrackingServerRef
type jsiiProxy_IMlflowTrackingServerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMlflowTrackingServerRef) MlflowTrackingServerRef() *MlflowTrackingServerReference {
	var returns *MlflowTrackingServerReference
	_jsii_.Get(
		j,
		"mlflowTrackingServerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMlflowTrackingServerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMlflowTrackingServerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

