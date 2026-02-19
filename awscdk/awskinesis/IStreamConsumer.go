package awskinesis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesis"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Kinesis Stream Consumer.
type IStreamConsumer interface {
	awscdk.IResource
	interfacesawskinesis.IStreamConsumerRef
	// Adds a statement to the IAM resource policy associated with this stream consumer.
	//
	// If this stream consumer was created in this stack (`new StreamConsumer`), a resource policy
	// will be automatically created upon the first call to `addToResourcePolicy`. If
	// the stream consumer is imported (`StreamConsumer.from`), then this is a no-op.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Grant the indicated permissions on this stream consumer to the provided IAM principal.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions for this stream consumer and its associated stream to an IAM principal (Role/Group/User).
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// The stream associated with this consumer.
	Stream() IStream
	// The ARN of the stream consumer.
	StreamConsumerArn() *string
	// The name of the stream consumer.
	StreamConsumerName() *string
}

// The jsii proxy for IStreamConsumer
type jsiiProxy_IStreamConsumer struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawskinesisIStreamConsumerRef
}

func (i *jsiiProxy_IStreamConsumer) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStreamConsumer) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IStreamConsumer) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IStreamConsumer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IStreamConsumer) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IStreamConsumer) Stream() IStream {
	var returns IStream
	_jsii_.Get(
		j,
		"stream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamConsumer) StreamConsumerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamConsumerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamConsumer) StreamConsumerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamConsumerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamConsumer) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamConsumer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamConsumer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamConsumer) StreamConsumerRef() *interfacesawskinesis.StreamConsumerReference {
	var returns *interfacesawskinesis.StreamConsumerReference
	_jsii_.Get(
		j,
		"streamConsumerRef",
		&returns,
	)
	return returns
}

