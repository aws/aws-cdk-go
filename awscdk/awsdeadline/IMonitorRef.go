package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Monitor.
// Experimental.
type IMonitorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMonitorRef
type jsiiProxy_IMonitorRef struct {
	internal.Type__constructsIConstruct
}

