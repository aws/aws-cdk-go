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

// Interface for GatewayTarget resources.
//
// Represents a target that hosts tools for the gateway.
// Targets can be Lambda functions, OpenAPI schemas, or Smithy models.
// Experimental.
type IGatewayTarget interface {
	interfacesawsbedrockagentcore.IGatewayTargetRef
	awscdk.IResource
	// Grants IAM actions to the IAM Principal.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grants `Create`, `Update`, and `Delete` actions on the Gateway Target.
	// Experimental.
	GrantManage(grantee awsiam.IGrantable) awsiam.Grant
	// Grants `Get` and `List` actions on the Gateway Target.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Timestamp when the gateway target was created.
	// Experimental.
	CreatedAt() *string
	// The credential provider configuration for the target.
	// Experimental.
	CredentialProviderConfigurations() *[]ICredentialProviderConfig
	// The description of the gateway target.
	// Experimental.
	Description() *string
	// The gateway that this target belongs to.
	// Experimental.
	Gateway() IGateway
	// The name of the gateway target.
	// Experimental.
	Name() *string
	// The status of the gateway target.
	// Experimental.
	Status() *string
	// The status reasons for the gateway target.
	// Experimental.
	StatusReasons() *[]*string
	// The ARN of the gateway target resource.
	// Experimental.
	TargetArn() *string
	// The id of the gateway target.
	// Experimental.
	TargetId() *string
	// The target protocol.
	// Experimental.
	TargetProtocolType() GatewayTargetProtocolType
	// Timestamp when the gateway target was last updated.
	// Experimental.
	UpdatedAt() *string
}

// The jsii proxy for IGatewayTarget
type jsiiProxy_IGatewayTarget struct {
	internal.Type__interfacesawsbedrockagentcoreIGatewayTargetRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGatewayTarget) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IGatewayTarget) GrantManage(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantManageParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantManage",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGatewayTarget) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IGatewayTarget) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IGatewayTarget) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) CredentialProviderConfigurations() *[]ICredentialProviderConfig {
	var returns *[]ICredentialProviderConfig
	_jsii_.Get(
		j,
		"credentialProviderConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) Gateway() IGateway {
	var returns IGateway
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) StatusReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) TargetProtocolType() GatewayTargetProtocolType {
	var returns GatewayTargetProtocolType
	_jsii_.Get(
		j,
		"targetProtocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) GatewayTargetRef() *interfacesawsbedrockagentcore.GatewayTargetReference {
	var returns *interfacesawsbedrockagentcore.GatewayTargetReference
	_jsii_.Get(
		j,
		"gatewayTargetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayTarget) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

