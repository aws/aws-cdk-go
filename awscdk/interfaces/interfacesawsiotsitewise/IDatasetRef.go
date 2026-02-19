package interfacesawsiotsitewise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Dataset.
// Experimental.
type IDatasetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Dataset resource.
	// Experimental.
	DatasetRef() *DatasetReference
}

// The jsii proxy for IDatasetRef
type jsiiProxy_IDatasetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDatasetRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDatasetRef) DatasetRef() *DatasetReference {
	var returns *DatasetReference
	_jsii_.Get(
		j,
		"datasetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatasetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatasetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

