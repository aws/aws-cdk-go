package awstransfer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebApp.
// Experimental.
type IWebAppRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a WebApp resource.
	// Experimental.
	WebAppRef() *WebAppReference
}

// The jsii proxy for IWebAppRef
type jsiiProxy_IWebAppRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IWebAppRef) WebAppRef() *WebAppReference {
	var returns *WebAppReference
	_jsii_.Get(
		j,
		"webAppRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebAppRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebAppRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

