package interfacesawsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OpsItem.
// Experimental.
type IOpsItemRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OpsItem resource.
	// Experimental.
	OpsItemRef() *OpsItemReference
}

// The jsii proxy for IOpsItemRef
type jsiiProxy_IOpsItemRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOpsItemRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOpsItemRef) OpsItemRef() *OpsItemReference {
	var returns *OpsItemReference
	_jsii_.Get(
		j,
		"opsItemRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOpsItemRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOpsItemRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

