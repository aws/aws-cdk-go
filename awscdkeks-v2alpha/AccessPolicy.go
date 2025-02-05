package awscdkeks-v2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeks-v2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an Amazon EKS Access Policy that implements the IAccessPolicy interface.
//
// Example:
//   // AmazonEKSClusterAdminPolicy with `cluster` scope
//   eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSClusterAdminPolicy"), &AccessPolicyNameOptions{
//   	AccessScopeType: eks.AccessScopeType_CLUSTER,
//   })
//   // AmazonEKSAdminPolicy with `namespace` scope
//   eks.AccessPolicy_FromAccessPolicyName(jsii.String("AmazonEKSAdminPolicy"), &AccessPolicyNameOptions{
//   	AccessScopeType: eks.AccessScopeType_NAMESPACE,
//   	Namespaces: []*string{
//   		jsii.String("foo"),
//   		jsii.String("bar"),
//   	},
//   })
//
// Experimental.
type AccessPolicy interface {
	IAccessPolicy
	// The scope of the access policy, which determines the level of access granted.
	// Experimental.
	AccessScope() *AccessScope
	// The access policy itself, which defines the specific permissions.
	// Experimental.
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
// Experimental.
func NewAccessPolicy(props *AccessPolicyProps) AccessPolicy {
	_init_.Initialize()

	if err := validateNewAccessPolicyParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPolicy{}

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicy",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AccessPolicy class.
// Experimental.
func NewAccessPolicy_Override(a AccessPolicy, props *AccessPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicy",
		[]interface{}{props},
		a,
	)
}

// Import AccessPolicy by name.
// Experimental.
func AccessPolicy_FromAccessPolicyName(policyName *string, options *AccessPolicyNameOptions) IAccessPolicy {
	_init_.Initialize()

	if err := validateAccessPolicy_FromAccessPolicyNameParameters(policyName, options); err != nil {
		panic(err)
	}
	var returns IAccessPolicy

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicy",
		"fromAccessPolicyName",
		[]interface{}{policyName, options},
		&returns,
	)

	return returns
}

