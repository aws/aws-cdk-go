package interfacesawswafv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswafv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebACL.
// Experimental.
type IWebACLRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WebACL resource.
	// Experimental.
	WebAclRef() *WebACLReference
}

// The jsii proxy for IWebACLRef
type jsiiProxy_IWebACLRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWebACLRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWebACLRef) WebAclRef() *WebACLReference {
	var returns *WebACLReference
	_jsii_.Get(
		j,
		"webAclRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebACLRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebACLRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

