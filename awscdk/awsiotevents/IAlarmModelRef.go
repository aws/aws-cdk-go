package awsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AlarmModel.
// Experimental.
type IAlarmModelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAlarmModelRef
type jsiiProxy_IAlarmModelRef struct {
	internal.Type__constructsIConstruct
}

