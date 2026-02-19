package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
	"github.com/aws/constructs-go/constructs/v10"
)

// An email identity.
type IEmailIdentity interface {
	interfacesawsses.IEmailIdentityRef
	awscdk.IResource
	// Adds an IAM policy statement associated with this email identity to an IAM principal's policy.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Permits an IAM principal the send email action.
	//
	// Actions: SendEmail.
	GrantSendEmail(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the email identity.
	EmailIdentityArn() *string
	// The name of the email identity.
	EmailIdentityName() *string
}

// The jsii proxy for IEmailIdentity
type jsiiProxy_IEmailIdentity struct {
	internal.Type__interfacesawssesIEmailIdentityRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEmailIdentity) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IEmailIdentity) GrantSendEmail(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantSendEmailParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantSendEmail",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEmailIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IEmailIdentity) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEmailIdentity) EmailIdentityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailIdentityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailIdentity) EmailIdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailIdentityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailIdentity) EmailIdentityRef() *interfacesawsses.EmailIdentityReference {
	var returns *interfacesawsses.EmailIdentityReference
	_jsii_.Get(
		j,
		"emailIdentityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailIdentity) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

