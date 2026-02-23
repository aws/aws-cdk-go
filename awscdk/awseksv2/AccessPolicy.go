package awseksv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an Amazon EKS Access Policy that implements the IAccessPolicy interface.
//
// Example:
//   var cluster Cluster
//   var nodeRole Role
//
//
//   // Grant access with EC2 type for Auto Mode node role
//   cluster.GrantAccess(jsii.String("nodeAccess"), nodeRole.roleArn, []IAccessPolicy{
//   	eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSAutoNodePolicy"), &AccessPolicyNameOptions{
//   		AccessScopeType: eks.AccessScopeType_CLUSTER,
//   	}),
//   }, &GrantAccessOptions{
//   	AccessEntryType: eks.AccessEntryType_EC2,
//   })
//
type AccessPolicy interface {
	IAccessPolicy
	// The scope of the access policy, which determines the level of access granted.
	AccessScope() *AccessScope
	// The access policy itself, which defines the specific permissions.
	Policy() *string
}

// The jsii proxy struct for AccessPolicy
type jsiiProxy_AccessPolicy struct {
	jsiiProxy_IAccessPolicy
}

func (j *jsiiProxy_AccessPolicy) AccessScope() *AccessScope {
	var returns *AccessScope
	_jsii_.Get(
		j,
		"accessScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}


// Constructs a new instance of the AccessPolicy class.
func NewAccessPolicy(props *AccessPolicyProps) AccessPolicy {
	_init_.Initialize()

	if err := validateNewAccessPolicyParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_eks_v2.AccessPolicy",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AccessPolicy class.
func NewAccessPolicy_Override(a AccessPolicy, props *AccessPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_eks_v2.AccessPolicy",
		[]interface{}{props},
		a,
	)
}

// Import AccessPolicy by name.
func AccessPolicy_FromAccessPolicyName(policyName *string, options *AccessPolicyNameOptions) IAccessPolicy {
	_init_.Initialize()

	if err := validateAccessPolicy_FromAccessPolicyNameParameters(policyName, options); err != nil {
		panic(err)
	}
	var returns IAccessPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks_v2.AccessPolicy",
		"fromAccessPolicyName",
		[]interface{}{policyName, options},
		&returns,
	)

	return returns
}

