package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Portal.
// Experimental.
type IPortalRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPortalRef
type jsiiProxy_IPortalRef struct {
	internal.Type__constructsIConstruct
}

