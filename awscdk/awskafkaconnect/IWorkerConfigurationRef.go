package awskafkaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkerConfiguration.
// Experimental.
type IWorkerConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a WorkerConfiguration resource.
	// Experimental.
	WorkerConfigurationRef() *WorkerConfigurationReference
}

// The jsii proxy for IWorkerConfigurationRef
type jsiiProxy_IWorkerConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IWorkerConfigurationRef) WorkerConfigurationRef() *WorkerConfigurationReference {
	var returns *WorkerConfigurationReference
	_jsii_.Get(
		j,
		"workerConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkerConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkerConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

