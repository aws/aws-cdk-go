package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HealthCheck.
// Experimental.
type IHealthCheckRef interface {
	constructs.IConstruct
	// A reference to a HealthCheck resource.
	// Experimental.
	HealthCheckRef() *HealthCheckReference
}

// The jsii proxy for IHealthCheckRef
type jsiiProxy_IHealthCheckRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHealthCheckRef) HealthCheckRef() *HealthCheckReference {
	var returns *HealthCheckReference
	_jsii_.Get(
		j,
		"healthCheckRef",
		&returns,
	)
	return returns
}

