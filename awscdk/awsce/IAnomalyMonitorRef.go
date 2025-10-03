package awsce

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalyMonitor.
// Experimental.
type IAnomalyMonitorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAnomalyMonitorRef
type jsiiProxy_IAnomalyMonitorRef struct {
	internal.Type__constructsIConstruct
}

