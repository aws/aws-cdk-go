package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Step Functions Activity https://docs.aws.amazon.com/step-functions/latest/dg/concepts-activities.html.
type IActivity interface {
	interfacesawsstepfunctions.IActivityRef
	awscdk.IResource
	// The ARN of the activity.
	ActivityArn() *string
	// The name of the activity.
	ActivityName() *string
	// The encryptionConfiguration object used for server-side encryption of the activity inputs.
	EncryptionConfiguration() EncryptionConfiguration
}

// The jsii proxy for IActivity
type jsiiProxy_IActivity struct {
	internal.Type__interfacesawsstepfunctionsIActivityRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IActivity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IActivity) ActivityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActivity) ActivityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActivity) EncryptionConfiguration() EncryptionConfiguration {
	var returns EncryptionConfiguration
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActivity) ActivityRef() *interfacesawsstepfunctions.ActivityReference {
	var returns *interfacesawsstepfunctions.ActivityReference
	_jsii_.Get(
		j,
		"activityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActivity) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActivity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActivity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

