package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam/internal"
)

// Represents an IAM Policy.
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage.html
//
// Experimental.
type IPolicy interface {
	awscdk.IResource
	// The name of this policy.
	// Experimental.
	PolicyName() *string
}

// The jsii proxy for IPolicy
type jsiiProxy_IPolicy struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPolicy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

