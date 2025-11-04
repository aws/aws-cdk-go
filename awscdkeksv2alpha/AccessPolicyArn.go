package awscdkeksv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an Amazon EKS Access Policy ARN.
//
// Amazon EKS Access Policies are used to control access to Amazon EKS clusters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   accessPolicyArn := eks_v2_alpha.AccessPolicyArn_AMAZON_EKS_ADMIN_POLICY()
//
// See: https://docs.aws.amazon.com/eks/latest/userguide/access-policies.html
//
// Experimental.
type AccessPolicyArn interface {
	// The Amazon Resource Name (ARN) of the access policy.
	// Experimental.
	PolicyArn() *string
	// - The name of the Amazon EKS access policy.
	//
	// This is used to construct the policy ARN.
	// Experimental.
	PolicyName() *string
}

// The jsii proxy struct for AccessPolicyArn
type jsiiProxy_AccessPolicyArn struct {
	_ byte // padding
}

func (j *jsiiProxy_AccessPolicyArn) PolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicyArn) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}


// Constructs a new instance of the `AccessEntry` class.
// Experimental.
func NewAccessPolicyArn(policyName *string) AccessPolicyArn {
	_init_.Initialize()

	if err := validateNewAccessPolicyArnParameters(policyName); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPolicyArn{}

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyArn",
		[]interface{}{policyName},
		&j,
	)

	return &j
}

// Constructs a new instance of the `AccessEntry` class.
// Experimental.
func NewAccessPolicyArn_Override(a AccessPolicyArn, policyName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyArn",
		[]interface{}{policyName},
		a,
	)
}

// Creates a new instance of the AccessPolicy class with the specified policy name.
//
// Returns: A new instance of the AccessPolicy class.
// Experimental.
func AccessPolicyArn_Of(policyName *string) AccessPolicyArn {
	_init_.Initialize()

	if err := validateAccessPolicyArn_OfParameters(policyName); err != nil {
		panic(err)
	}
	var returns AccessPolicyArn

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyArn",
		"of",
		[]interface{}{policyName},
		&returns,
	)

	return returns
}

func AccessPolicyArn_AMAZON_EKS_ADMIN_POLICY() AccessPolicyArn {
	_init_.Initialize()
	var returns AccessPolicyArn
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyArn",
		"AMAZON_EKS_ADMIN_POLICY",
		&returns,
	)
	return returns
}

func AccessPolicyArn_AMAZON_EKS_ADMIN_VIEW_POLICY() AccessPolicyArn {
	_init_.Initialize()
	var returns AccessPolicyArn
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyArn",
		"AMAZON_EKS_ADMIN_VIEW_POLICY",
		&returns,
	)
	return returns
}

func AccessPolicyArn_AMAZON_EKS_CLUSTER_ADMIN_POLICY() AccessPolicyArn {
	_init_.Initialize()
	var returns AccessPolicyArn
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyArn",
		"AMAZON_EKS_CLUSTER_ADMIN_POLICY",
		&returns,
	)
	return returns
}

func AccessPolicyArn_AMAZON_EKS_EDIT_POLICY() AccessPolicyArn {
	_init_.Initialize()
	var returns AccessPolicyArn
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyArn",
		"AMAZON_EKS_EDIT_POLICY",
		&returns,
	)
	return returns
}

func AccessPolicyArn_AMAZON_EKS_VIEW_POLICY() AccessPolicyArn {
	_init_.Initialize()
	var returns AccessPolicyArn
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyArn",
		"AMAZON_EKS_VIEW_POLICY",
		&returns,
	)
	return returns
}

