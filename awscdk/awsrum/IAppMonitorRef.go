package awsrum

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrum/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppMonitor.
// Experimental.
type IAppMonitorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAppMonitorRef
type jsiiProxy_IAppMonitorRef struct {
	internal.Type__constructsIConstruct
}

