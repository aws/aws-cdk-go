package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswaf/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebACL.
// Experimental.
type IWebACLRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a WebACL resource.
	// Experimental.
	WebAclRef() *WebACLReference
}

// The jsii proxy for IWebACLRef
type jsiiProxy_IWebACLRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IWebACLRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

