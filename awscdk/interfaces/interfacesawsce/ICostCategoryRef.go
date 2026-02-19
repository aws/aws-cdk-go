package interfacesawsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CostCategory.
// Experimental.
type ICostCategoryRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CostCategory resource.
	// Experimental.
	CostCategoryRef() *CostCategoryReference
}

// The jsii proxy for ICostCategoryRef
type jsiiProxy_ICostCategoryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICostCategoryRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICostCategoryRef) CostCategoryRef() *CostCategoryReference {
	var returns *CostCategoryReference
	_jsii_.Get(
		j,
		"costCategoryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICostCategoryRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICostCategoryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

