package interfacesawsobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsobservabilityadmin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DatasetIntegration.
// Experimental.
type IDatasetIntegrationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DatasetIntegration resource.
	// Experimental.
	DatasetIntegrationRef() *DatasetIntegrationReference
}

// The jsii proxy for IDatasetIntegrationRef
type jsiiProxy_IDatasetIntegrationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDatasetIntegrationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDatasetIntegrationRef) DatasetIntegrationRef() *DatasetIntegrationReference {
	var returns *DatasetIntegrationReference
	_jsii_.Get(
		j,
		"datasetIntegrationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatasetIntegrationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatasetIntegrationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

