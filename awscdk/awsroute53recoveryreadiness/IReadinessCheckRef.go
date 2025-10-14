package awsroute53recoveryreadiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoveryreadiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReadinessCheck.
// Experimental.
type IReadinessCheckRef interface {
	constructs.IConstruct
	// A reference to a ReadinessCheck resource.
	// Experimental.
	ReadinessCheckRef() *ReadinessCheckReference
}

// The jsii proxy for IReadinessCheckRef
type jsiiProxy_IReadinessCheckRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IReadinessCheckRef) ReadinessCheckRef() *ReadinessCheckReference {
	var returns *ReadinessCheckReference
	_jsii_.Get(
		j,
		"readinessCheckRef",
		&returns,
	)
	return returns
}

