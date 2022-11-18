// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Your custom-built game server software that runs on GameLift and hosts game sessions for your players.
//
// A game build represents the set of files that run your game server on a particular operating system.
// You can have many different builds, such as for different flavors of your game.
// The game build must be integrated with the GameLift service.
// You upload game build files to the GameLift service in the Regions where you plan to set up fleets.
// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-cli-uploading.html
//
// Experimental.
type IBuild interface {
	awsiam.IGrantable
	awscdk.IResource
	// The Identifier of the build.
	// Experimental.
	BuildId() *string
}

// The jsii proxy for IBuild
type jsiiProxy_IBuild struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBuild) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IBuild) BuildId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuild) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuild) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuild) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuild) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

