package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RealtimeLogConfig.
// Experimental.
type IRealtimeLogConfigRef interface {
	constructs.IConstruct
	// A reference to a RealtimeLogConfig resource.
	// Experimental.
	RealtimeLogConfigRef() *RealtimeLogConfigReference
}

// The jsii proxy for IRealtimeLogConfigRef
type jsiiProxy_IRealtimeLogConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRealtimeLogConfigRef) RealtimeLogConfigRef() *RealtimeLogConfigReference {
	var returns *RealtimeLogConfigReference
	_jsii_.Get(
		j,
		"realtimeLogConfigRef",
		&returns,
	)
	return returns
}

