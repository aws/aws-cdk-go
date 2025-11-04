package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for Memory resources.
// Experimental.
type IMemory interface {
	awsiam.IGrantable
	awscdk.IResource
	// Grant the given principal identity permissions to perform actions on this memory.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given principal identity permissions to manage the control plane of this memory.
	// Experimental.
	GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given principal identity permissions to delete content on this memory.
	// Experimental.
	GrantDelete(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given principal identity permissions to delete Long-Term Memory (LTM) content on this memory.
	// Experimental.
	GrantDeleteLongTermMemory(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given principal identity permissions to delete Short-Term Memory (STM) content on this memory.
	// Experimental.
	GrantDeleteShortTermMemory(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given principal identity permissions to do every action on this memory.
	// Experimental.
	GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given principal identity permissions to read the contents of this memory.
	//
	// Both Short-Term Memory (STM) and Long-Term Memory (LTM).
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given principal identity permissions to read the Long-Term Memory (LTM) contents of this memory.
	// Experimental.
	GrantReadLongTermMemory(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given principal identity permissions to read the Short-Term Memory (STM) contents of this memory.
	// Experimental.
	GrantReadShortTermMemory(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given principal identity permissions to write content to this memory.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this memory.
	// Experimental.
	Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of errors for a specific API operation performed on this memory.
	// Experimental.
	MetricErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns the metric containing the number of created memory events and memory records.
	// Experimental.
	MetricEventCreationCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the given named metric related to the API operation performed on this memory.
	// Experimental.
	MetricForApiOperation(metricName *string, operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of API requests made for a specific memory operation.
	// Experimental.
	MetricInvocationsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the latency of a specific API operation performed on this memory.
	// Experimental.
	MetricLatencyForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Timestamp when the memory was created.
	// Experimental.
	CreatedAt() *string
	// The IAM role that provides permissions for the memory to access AWS services.
	// Experimental.
	ExecutionRole() awsiam.IRole
	// Custom KMS key for encryption (if provided).
	// Experimental.
	KmsKey() awskms.IKey
	// The ARN of the memory resource.
	// Experimental.
	MemoryArn() *string
	// The id of the memory.
	// Experimental.
	MemoryId() *string
	// The status of the memory.
	// Experimental.
	Status() *string
	// Timestamp when the memory was last updated.
	// Experimental.
	UpdatedAt() *string
}

// The jsii proxy for IMemory
type jsiiProxy_IMemory struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IMemory) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IMemory) GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantAdminParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantAdmin",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) GrantDelete(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantDeleteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDelete",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) GrantDeleteLongTermMemory(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantDeleteLongTermMemoryParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDeleteLongTermMemory",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) GrantDeleteShortTermMemory(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantDeleteShortTermMemoryParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDeleteShortTermMemory",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantFullAccessParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantFullAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) GrantReadLongTermMemory(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadLongTermMemoryParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadLongTermMemory",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) GrantReadShortTermMemory(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadShortTermMemoryParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadShortTermMemory",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, dimensions, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, dimensions, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) MetricErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricErrorsForApiOperationParameters(operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricErrorsForApiOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) MetricEventCreationCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEventCreationCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEventCreationCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) MetricForApiOperation(metricName *string, operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricForApiOperationParameters(metricName, operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricForApiOperation",
		[]interface{}{metricName, operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) MetricInvocationsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationsForApiOperationParameters(operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationsForApiOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) MetricLatencyForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricLatencyForApiOperationParameters(operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricLatencyForApiOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IMemory) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IMemory) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) MemoryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) MemoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemory) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

