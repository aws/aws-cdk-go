package awsbedrockalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents an Agent Alias, either created with CDK or imported.
// Experimental.
type IAgentAlias interface {
	awscdk.IResource
	// Grant the given principal identity permissions to perform actions on this agent alias.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to get the agent alias.
	// Experimental.
	GrantGet(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to invoke the agent alias.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Define an EventBridge rule that triggers when something happens to this agent alias.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	// Experimental.
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The underlying agent for this alias.
	// Experimental.
	Agent() IAgent
	// The ARN of the agent alias.
	// Experimental.
	AliasArn() *string
	// The unique identifier of the agent alias.
	// Experimental.
	AliasId() *string
}

// The jsii proxy for IAgentAlias
type jsiiProxy_IAgentAlias struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAgentAlias) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IAgentAlias) GrantGet(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantGetParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantGet",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAgentAlias) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAgentAlias) OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnCloudTrailEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAgentAlias) Agent() IAgent {
	var returns IAgent
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentAlias) AliasArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentAlias) AliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasId",
		&returns,
	)
	return returns
}

