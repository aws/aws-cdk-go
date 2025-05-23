package awseks


// An access policy includes permissions that allow Amazon EKS to authorize an IAM principal to work with Kubernetes objects on your cluster.
//
// The policies are managed by Amazon EKS, but they're not IAM policies. You can't view the permissions in the policies using the API. The permissions for many of the policies are similar to the Kubernetes `cluster-admin` , `admin` , `edit` , and `view` cluster roles. For more information about these cluster roles, see [User-facing roles](https://docs.aws.amazon.com/https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles) in the Kubernetes documentation. To view the contents of the policies, see [Access policy permissions](https://docs.aws.amazon.com/eks/latest/userguide/access-policies.html#access-policy-permissions) in the *Amazon EKS User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPolicyProperty := &AccessPolicyProperty{
//   	AccessScope: &AccessScopeProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   	},
//   	PolicyArn: jsii.String("policyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-accessentry-accesspolicy.html
//
type CfnAccessEntry_AccessPolicyProperty struct {
	// The scope of an `AccessPolicy` that's associated to an `AccessEntry` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-accessentry-accesspolicy.html#cfn-eks-accessentry-accesspolicy-accessscope
	//
	AccessScope interface{} `field:"required" json:"accessScope" yaml:"accessScope"`
	// The ARN of the access policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-accessentry-accesspolicy.html#cfn-eks-accessentry-accesspolicy-policyarn
	//
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
}

