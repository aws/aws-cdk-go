// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/internal"
)

// Represents a Gamelift Alias for a Gamelift fleet destination.
// Experimental.
type IAlias interface {
	awscdk.IResource
	// The ARN of the alias.
	// Experimental.
	AliasArn() *string
	// The Identifier of the alias.
	// Experimental.
	AliasId() *string
}

// The jsii proxy for IAlias
type jsiiProxy_IAlias struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAlias) AliasArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) AliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasId",
		&returns,
	)
	return returns
}

