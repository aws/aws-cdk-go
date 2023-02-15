// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcloud9alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// An environment owner.
//
// Example:
//   var vpc vpc
//
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	imageId: cloud9.imageId_AMAZON_LINUX_2,
//
//   	owner: cloud9.owner.accountRoot(jsii.String("111111111")),
//   })
//
// Experimental.
type Owner interface {
	// of environment owner.
	// Experimental.
	OwnerArn() *string
}

// The jsii proxy struct for Owner
type jsiiProxy_Owner struct {
	_ byte // padding
}

func (j *jsiiProxy_Owner) OwnerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerArn",
		&returns,
	)
	return returns
}


// Make the Account Root User the environment owner (not recommended).
// Experimental.
func Owner_AccountRoot(accountId *string) Owner {
	_init_.Initialize()

	if err := validateOwner_AccountRootParameters(accountId); err != nil {
		panic(err)
	}
	var returns Owner

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.Owner",
		"accountRoot",
		[]interface{}{accountId},
		&returns,
	)

	return returns
}

// Make an IAM user the environment owner.
//
// User need to have AWSCloud9Administrator permissions.
// See: https://docs.aws.amazon.com/cloud9/latest/user-guide/share-environment.html#share-environment-about
//
// Experimental.
func Owner_User(user awsiam.IUser) Owner {
	_init_.Initialize()

	if err := validateOwner_UserParameters(user); err != nil {
		panic(err)
	}
	var returns Owner

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.Owner",
		"user",
		[]interface{}{user},
		&returns,
	)

	return returns
}

