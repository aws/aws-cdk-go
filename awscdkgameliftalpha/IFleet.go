// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Gamelift fleet.
// Experimental.
type IFleet interface {
	IGameSessionQueueDestination
	awsiam.IGrantable
	awscdk.IResource
	// Grant the `grantee` identity permissions to perform `actions`.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Return the given named metric for this fleet.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Instances with `ACTIVE` status, which means they are running active server processes.
	//
	// The count includes idle instances and those that are hosting one or more game sessions.
	// This metric measures current total instance capacity.
	//
	// This metric can be used with automatic scaling.
	// Experimental.
	MetricActiveInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Target number of active instances that GameLift is working to maintain in the fleet.
	//
	// With automatic scaling, this value is determined based on the scaling policies currently in force.
	// Without automatic scaling, this value is set manually.
	// This metric is not available when viewing data for fleet metric groups.
	// Experimental.
	MetricDesiredInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Active instances that are currently hosting zero (0) game sessions.
	//
	// This metric measures capacity that is available but unused.
	// This metric can be used with automatic scaling.
	// Experimental.
	MetricIdleInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Number of spot instances that have been interrupted.
	// Experimental.
	MetricInstanceInterruptions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Maximum number of instances that are allowed for the fleet.
	//
	// A fleet's instance maximum determines the capacity ceiling during manual or automatic scaling up.
	// This metric is not available when viewing data for fleet metric groups.
	// Experimental.
	MetricMaxInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Minimum number of instances allowed for the fleet.
	//
	// A fleet's instance minimum determines the capacity floor during manual or automatic scaling down.
	// This metric is not available when viewing data for fleet metric groups.
	// Experimental.
	MetricMinInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Percentage of all active instances that are idle (calculated as IdleInstances / ActiveInstances).
	//
	// This metric can be used for automatic scaling.
	// Experimental.
	MetricPercentIdleInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the fleet.
	// Experimental.
	FleetArn() *string
	// The Identifier of the fleet.
	// Experimental.
	FleetId() *string
}

// The jsii proxy for IFleet
type jsiiProxy_IFleet struct {
	jsiiProxy_IGameSessionQueueDestination
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFleet) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFleet) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFleet) MetricActiveInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricActiveInstancesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricActiveInstances",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFleet) MetricDesiredInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDesiredInstancesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDesiredInstances",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFleet) MetricIdleInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIdleInstancesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIdleInstances",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFleet) MetricInstanceInterruptions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInstanceInterruptionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInstanceInterruptions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFleet) MetricMaxInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMaxInstancesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMaxInstances",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFleet) MetricMinInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricMinInstancesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricMinInstances",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFleet) MetricPercentIdleInstances(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPercentIdleInstancesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPercentIdleInstances",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFleet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IFleet) FleetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) FleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) ResourceArnForDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnForDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFleet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

