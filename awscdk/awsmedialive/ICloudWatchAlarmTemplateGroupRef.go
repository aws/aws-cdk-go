package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudWatchAlarmTemplateGroup.
// Experimental.
type ICloudWatchAlarmTemplateGroupRef interface {
	constructs.IConstruct
	// A reference to a CloudWatchAlarmTemplateGroup resource.
	// Experimental.
	CloudWatchAlarmTemplateGroupRef() *CloudWatchAlarmTemplateGroupReference
}

// The jsii proxy for ICloudWatchAlarmTemplateGroupRef
type jsiiProxy_ICloudWatchAlarmTemplateGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICloudWatchAlarmTemplateGroupRef) CloudWatchAlarmTemplateGroupRef() *CloudWatchAlarmTemplateGroupReference {
	var returns *CloudWatchAlarmTemplateGroupReference
	_jsii_.Get(
		j,
		"cloudWatchAlarmTemplateGroupRef",
		&returns,
	)
	return returns
}

