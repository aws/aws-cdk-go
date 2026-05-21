package interfacesawsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Browser.
// Experimental.
type IBrowserRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Browser resource.
	// Experimental.
	BrowserRef() *BrowserReference
}

// The jsii proxy for IBrowserRef
type jsiiProxy_IBrowserRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IBrowserRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBrowserRef) BrowserRef() *BrowserReference {
	var returns *BrowserReference
	_jsii_.Get(
		j,
		"browserRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

