package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebApp.
// Experimental.
type IWebAppRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWebAppRef
type jsiiProxy_IWebAppRef struct {
	internal.Type__constructsIConstruct
}

