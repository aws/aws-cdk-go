package interfacesawsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataProvider.
// Experimental.
type IDataProviderRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataProvider resource.
	// Experimental.
	DataProviderRef() *DataProviderReference
}

// The jsii proxy for IDataProviderRef
type jsiiProxy_IDataProviderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDataProviderRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDataProviderRef) DataProviderRef() *DataProviderReference {
	var returns *DataProviderReference
	_jsii_.Get(
		j,
		"dataProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataProviderRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataProviderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

