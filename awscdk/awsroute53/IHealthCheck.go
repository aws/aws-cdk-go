package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
)

// Imported or created health check.
type IHealthCheck interface {
	awscdk.IResource
	// The ID of the health check.
	HealthCheckId() *string
}

// The jsii proxy for IHealthCheck
type jsiiProxy_IHealthCheck struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IHealthCheck) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

