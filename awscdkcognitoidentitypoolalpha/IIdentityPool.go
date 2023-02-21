// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha/v2/internal"
)

// Represents a Cognito IdentityPool.
// Experimental.
type IIdentityPool interface {
	awscdk.IResource
	// The ARN of the Identity Pool.
	// Experimental.
	IdentityPoolArn() *string
	// The id of the Identity Pool in the format REGION:GUID.
	// Experimental.
	IdentityPoolId() *string
	// Name of the Identity Pool.
	// Experimental.
	IdentityPoolName() *string
}

// The jsii proxy for IIdentityPool
type jsiiProxy_IIdentityPool struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IIdentityPool) IdentityPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) IdentityPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolName",
		&returns,
	)
	return returns
}

