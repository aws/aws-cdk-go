package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CompositeAlarm.
// Experimental.
type ICompositeAlarmRef interface {
	constructs.IConstruct
	// A reference to a CompositeAlarm resource.
	// Experimental.
	CompositeAlarmRef() *CompositeAlarmReference
}

// The jsii proxy for ICompositeAlarmRef
type jsiiProxy_ICompositeAlarmRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICompositeAlarmRef) CompositeAlarmRef() *CompositeAlarmReference {
	var returns *CompositeAlarmReference
	_jsii_.Get(
		j,
		"compositeAlarmRef",
		&returns,
	)
	return returns
}

