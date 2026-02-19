package interfacesawsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaMapping.
// Experimental.
type ISchemaMappingRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SchemaMapping resource.
	// Experimental.
	SchemaMappingRef() *SchemaMappingReference
}

// The jsii proxy for ISchemaMappingRef
type jsiiProxy_ISchemaMappingRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISchemaMappingRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISchemaMappingRef) SchemaMappingRef() *SchemaMappingReference {
	var returns *SchemaMappingReference
	_jsii_.Get(
		j,
		"schemaMappingRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchemaMappingRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchemaMappingRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

