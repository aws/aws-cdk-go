package interfacesawsdevopsguru

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdevopsguru/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogAnomalyDetectionIntegration.
// Experimental.
type ILogAnomalyDetectionIntegrationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LogAnomalyDetectionIntegration resource.
	// Experimental.
	LogAnomalyDetectionIntegrationRef() *LogAnomalyDetectionIntegrationReference
}

// The jsii proxy for ILogAnomalyDetectionIntegrationRef
type jsiiProxy_ILogAnomalyDetectionIntegrationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILogAnomalyDetectionIntegrationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ILogAnomalyDetectionIntegrationRef) LogAnomalyDetectionIntegrationRef() *LogAnomalyDetectionIntegrationReference {
	var returns *LogAnomalyDetectionIntegrationReference
	_jsii_.Get(
		j,
		"logAnomalyDetectionIntegrationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILogAnomalyDetectionIntegrationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILogAnomalyDetectionIntegrationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

