package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Catalog.
// Experimental.
type ICatalogRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Catalog resource.
	// Experimental.
	CatalogRef() *CatalogReference
}

// The jsii proxy for ICatalogRef
type jsiiProxy_ICatalogRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICatalogRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICatalogRef) CatalogRef() *CatalogReference {
	var returns *CatalogReference
	_jsii_.Get(
		j,
		"catalogRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICatalogRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICatalogRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

