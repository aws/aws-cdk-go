package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudWatchAlarmTemplate.
// Experimental.
type ICloudWatchAlarmTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICloudWatchAlarmTemplateRef
type jsiiProxy_ICloudWatchAlarmTemplateRef struct {
	internal.Type__constructsIConstruct
}

