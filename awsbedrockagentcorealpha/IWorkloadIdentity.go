package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore"
	"github.com/aws/constructs-go/constructs/v10"
)

// A workload identity for Amazon Bedrock AgentCore.
//
// Represents the stable identity of an agent within an account's agent identity directory.
// It ties together IAM roles, OAuth2 flows, API keys, and workload access tokens
// for consistent authentication across environments.
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/understanding-agent-identities.html
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type IWorkloadIdentity interface {
	awsiam.IGrantable
	awscdk.IResource
	interfacesawsbedrockagentcore.IWorkloadIdentityRef
	// Grants IAM actions on this workload identity, scoped to its ARN and the parent resources required by the Bedrock AgentCore authorization model.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant control plane permissions to manage this workload identity.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant
	// Grant read, list, admin, and use permissions.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant
	// Grant `GetWorkloadIdentity` and `ListWorkloadIdentities`, scoped to this identity and parent resources required by the Bedrock AgentCore authorization model.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grant data plane permissions to mint workload access tokens (`GetWorkloadAccessToken`, `GetWorkloadAccessTokenForJWT`, `GetWorkloadAccessTokenForUserId`).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GrantUse(grantee awsiam.IGrantable) awsiam.Grant
	// Timestamp when the workload identity was created.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CreatedTime() *string
	// Timestamp when the workload identity was last updated.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LastUpdatedTime() *string
	// The ARN of this workload identity.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	WorkloadIdentityArn() *string
	// The name of this workload identity.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	WorkloadIdentityName() *string
}

// The jsii proxy for IWorkloadIdentity
type jsiiProxy_IWorkloadIdentity struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
	internal.Type__interfacesawsbedrockagentcoreIWorkloadIdentityRef
}

func (i *jsiiProxy_IWorkloadIdentity) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IWorkloadIdentity) GrantAdmin(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IWorkloadIdentity) GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IWorkloadIdentity) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IWorkloadIdentity) GrantUse(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantUseParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantUse",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWorkloadIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IWorkloadIdentity) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWorkloadIdentity) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentity) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentity) WorkloadIdentityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadIdentityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentity) WorkloadIdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadIdentityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentity) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentity) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentity) WorkloadIdentityRef() *interfacesawsbedrockagentcore.WorkloadIdentityReference {
	var returns *interfacesawsbedrockagentcore.WorkloadIdentityReference
	_jsii_.Get(
		j,
		"workloadIdentityRef",
		&returns,
	)
	return returns
}

