// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Your configuration and custom game logic for use with Realtime Servers.
//
// Realtime Servers are provided by GameLift to use instead of a custom-built game server.
// You configure Realtime Servers for your game clients by creating a script using JavaScript,
// and add custom game logic as appropriate to host game sessions for your players.
// You upload the Realtime script to the GameLift service in the Regions where you plan to set up fleets.
// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-script-uploading.html
//
// Experimental.
type IScript interface {
	awsiam.IGrantable
	awscdk.IResource
	// The ARN of the realtime server script.
	// Experimental.
	ScriptArn() *string
	// The Identifier of the realtime server script.
	// Experimental.
	ScriptId() *string
}

// The jsii proxy for IScript
type jsiiProxy_IScript struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IScript) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IScript) ScriptArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScript) ScriptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScript) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScript) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScript) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScript) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

