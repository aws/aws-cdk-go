package awstransfer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebApp.
// Experimental.
type IWebAppRef interface {
	constructs.IConstruct
	// A reference to a WebApp resource.
	// Experimental.
	WebAppRef() *WebAppReference
}

// The jsii proxy for IWebAppRef
type jsiiProxy_IWebAppRef struct {
	internal.Type__constructsIConstruct
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

