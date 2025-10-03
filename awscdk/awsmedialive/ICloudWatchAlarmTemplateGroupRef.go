package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudWatchAlarmTemplateGroup.
// Experimental.
type ICloudWatchAlarmTemplateGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICloudWatchAlarmTemplateGroupRef
type jsiiProxy_ICloudWatchAlarmTemplateGroupRef struct {
	internal.Type__constructsIConstruct
}

