package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAccessEntry`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessEntryProps := &CfnAccessEntryProps{
//   	ClusterName: jsii.String("clusterName"),
//   	PrincipalArn: jsii.String("principalArn"),
//
//   	// the properties below are optional
//   	AccessPolicies: []interface{}{
//   		&AccessPolicyProperty{
//   			AccessScope: &AccessScopeProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   			},
//   			PolicyArn: jsii.String("policyArn"),
//   		},
//   	},
//   	KubernetesGroups: []*string{
//   		jsii.String("kubernetesGroups"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-accessentry.html
//
type CfnAccessEntryProps struct {
	// The name of your cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-accessentry.html#cfn-eks-accessentry-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The ARN of the IAM principal for the `AccessEntry` .
	//
	// You can specify one ARN for each access entry. You can't specify the same ARN in more than one access entry. This value can't be changed after access entry creation.
	//
	// The valid principals differ depending on the type of the access entry in the `type` field. The only valid ARN is IAM roles for the types of access entries for nodes: `` `` . You can use every IAM principal type for `STANDARD` access entries. You can't use the STS session principal type with access entries because this is a temporary principal for each session and not a permanent identity that can be assigned permissions.
	//
	// [IAM best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#bp-users-federation-idp) recommend using IAM roles with temporary credentials, rather than IAM users with long-term credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-accessentry.html#cfn-eks-accessentry-principalarn
	//
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
	// The access policies to associate to the access entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-accessentry.html#cfn-eks-accessentry-accesspolicies
	//
	AccessPolicies interface{} `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// The value for `name` that you've specified for `kind: Group` as a `subject` in a Kubernetes `RoleBinding` or `ClusterRoleBinding` object.
	//
	// Amazon EKS doesn't confirm that the value for `name` exists in any bindings on your cluster. You can specify one or more names.
	//
	// Kubernetes authorizes the `principalArn` of the access entry to access any cluster objects that you've specified in a Kubernetes `Role` or `ClusterRole` object that is also specified in a binding's `roleRef` . For more information about creating Kubernetes `RoleBinding` , `ClusterRoleBinding` , `Role` , or `ClusterRole` objects, see [Using RBAC Authorization in the Kubernetes documentation](https://docs.aws.amazon.com/https://kubernetes.io/docs/reference/access-authn-authz/rbac/) .
	//
	// If you want Amazon EKS to authorize the `principalArn` (instead of, or in addition to Kubernetes authorizing the `principalArn` ), you can associate one or more access policies to the access entry using `AssociateAccessPolicy` . If you associate any access policies, the `principalARN` has all permissions assigned in the associated access policies and all permissions in any Kubernetes `Role` or `ClusterRole` objects that the group names are bound to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-accessentry.html#cfn-eks-accessentry-kubernetesgroups
	//
	KubernetesGroups *[]*string `field:"optional" json:"kubernetesGroups" yaml:"kubernetesGroups"`
	// Metadata that assists with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Tags don't propagate to any other cluster or AWS resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-accessentry.html#cfn-eks-accessentry-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the new access entry. Valid values are `Standard` , `FARGATE_LINUX` , `EC2_LINUX` , and `EC2_WINDOWS` .
	//
	// If the `principalArn` is for an IAM role that's used for self-managed Amazon EC2 nodes, specify `EC2_LINUX` or `EC2_WINDOWS` . Amazon EKS grants the necessary permissions to the node for you. If the `principalArn` is for any other purpose, specify `STANDARD` . If you don't specify a value, Amazon EKS sets the value to `STANDARD` . It's unnecessary to create access entries for IAM roles used with Fargate profiles or managed Amazon EC2 nodes, because Amazon EKS creates entries in the `aws-auth` `ConfigMap` for the roles. You can't change this value once you've created the access entry.
	//
	// If you set the value to `EC2_LINUX` or `EC2_WINDOWS` , you can't specify values for `kubernetesGroups` , or associate an `AccessPolicy` to the access entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-accessentry.html#cfn-eks-accessentry-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The username to authenticate to Kubernetes with.
	//
	// We recommend not specifying a username and letting Amazon EKS specify it for you. For more information about the value Amazon EKS specifies for you, or constraints before specifying your own username, see [Creating access entries](https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html#creating-access-entries) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-accessentry.html#cfn-eks-accessentry-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

