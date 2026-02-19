package interfacesawsforecast

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsforecast/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DatasetGroup.
// Experimental.
type IDatasetGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DatasetGroup resource.
	// Experimental.
	DatasetGroupRef() *DatasetGroupReference
}

// The jsii proxy for IDatasetGroupRef
type jsiiProxy_IDatasetGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDatasetGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDatasetGroupRef) DatasetGroupRef() *DatasetGroupReference {
	var returns *DatasetGroupReference
	_jsii_.Get(
		j,
		"datasetGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatasetGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatasetGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

