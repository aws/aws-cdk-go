package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelPackageGroup.
// Experimental.
type IModelPackageGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ModelPackageGroup resource.
	// Experimental.
	ModelPackageGroupRef() *ModelPackageGroupReference
}

// The jsii proxy for IModelPackageGroupRef
type jsiiProxy_IModelPackageGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IModelPackageGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IModelPackageGroupRef) ModelPackageGroupRef() *ModelPackageGroupReference {
	var returns *ModelPackageGroupReference
	_jsii_.Get(
		j,
		"modelPackageGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelPackageGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelPackageGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

