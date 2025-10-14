package awsshield

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProactiveEngagement.
// Experimental.
type IProactiveEngagementRef interface {
	constructs.IConstruct
	// A reference to a ProactiveEngagement resource.
	// Experimental.
	ProactiveEngagementRef() *ProactiveEngagementReference
}

// The jsii proxy for IProactiveEngagementRef
type jsiiProxy_IProactiveEngagementRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProactiveEngagementRef) ProactiveEngagementRef() *ProactiveEngagementReference {
	var returns *ProactiveEngagementReference
	_jsii_.Get(
		j,
		"proactiveEngagementRef",
		&returns,
	)
	return returns
}

