package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SessionLogger.
// Experimental.
type ISessionLoggerRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISessionLoggerRef
type jsiiProxy_ISessionLoggerRef struct {
	internal.Type__constructsIConstruct
}

