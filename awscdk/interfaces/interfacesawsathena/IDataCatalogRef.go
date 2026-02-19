package interfacesawsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataCatalog.
// Experimental.
type IDataCatalogRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataCatalog resource.
	// Experimental.
	DataCatalogRef() *DataCatalogReference
}

// The jsii proxy for IDataCatalogRef
type jsiiProxy_IDataCatalogRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDataCatalogRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDataCatalogRef) DataCatalogRef() *DataCatalogReference {
	var returns *DataCatalogReference
	_jsii_.Get(
		j,
		"dataCatalogRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataCatalogRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataCatalogRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

