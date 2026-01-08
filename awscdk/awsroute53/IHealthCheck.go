package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

// Imported or created health check.
type IHealthCheck interface {
	interfacesawsroute53.IHealthCheckRef
	awscdk.IResource
	// The ID of the health check.
	HealthCheckId() *string
}

// The jsii proxy for IHealthCheck
type jsiiProxy_IHealthCheck struct {
	internal.Type__interfacesawsroute53IHealthCheckRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IHealthCheck) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IHealthCheck) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHealthCheck) HealthCheckRef() *interfacesawsroute53.HealthCheckReference {
	var returns *interfacesawsroute53.HealthCheckReference
	_jsii_.Get(
		j,
		"healthCheckRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHealthCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHealthCheck) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

