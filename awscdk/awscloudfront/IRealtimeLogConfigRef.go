package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RealtimeLogConfig.
// Experimental.
type IRealtimeLogConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRealtimeLogConfigRef
type jsiiProxy_IRealtimeLogConfigRef struct {
	internal.Type__constructsIConstruct
}

