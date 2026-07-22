package interfacesawspersonalize

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataDeletionJob.
// Experimental.
type IDataDeletionJobRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataDeletionJob resource.
	// Experimental.
	DataDeletionJobRef() *DataDeletionJobReference
}

// The jsii proxy for IDataDeletionJobRef
type jsiiProxy_IDataDeletionJobRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDataDeletionJobRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDataDeletionJobRef) DataDeletionJobRef() *DataDeletionJobReference {
	var returns *DataDeletionJobReference
	_jsii_.Get(
		j,
		"dataDeletionJobRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataDeletionJobRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataDeletionJobRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

