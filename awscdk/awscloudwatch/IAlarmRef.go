package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Alarm.
// Experimental.
type IAlarmRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAlarmRef
type jsiiProxy_IAlarmRef struct {
	internal.Type__constructsIConstruct
}

