package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPodIdentityAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPodIdentityAssociationProps := &CfnPodIdentityAssociationProps{
//   	ClusterName: jsii.String("clusterName"),
//   	Namespace: jsii.String("namespace"),
//   	RoleArn: jsii.String("roleArn"),
//   	ServiceAccount: jsii.String("serviceAccount"),
//
//   	// the properties below are optional
//   	DisableSessionTags: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetRoleArn: jsii.String("targetRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-podidentityassociation.html
//
type CfnPodIdentityAssociationProps struct {
	// The name of the cluster that the association is in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-podidentityassociation.html#cfn-eks-podidentityassociation-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The name of the Kubernetes namespace inside the cluster to create the association in.
	//
	// The service account and the Pods that use the service account must be in this namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-podidentityassociation.html#cfn-eks-podidentityassociation-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with the service account.
	//
	// The EKS Pod Identity agent manages credentials to assume this role for applications in the containers in the Pods that use this service account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-podidentityassociation.html#cfn-eks-podidentityassociation-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the Kubernetes service account inside the cluster to associate the IAM credentials with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-podidentityassociation.html#cfn-eks-podidentityassociation-serviceaccount
	//
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// The state of the automatic sessions tags. The value of *true* disables these tags.
	//
	// EKS Pod Identity adds a pre-defined set of session tags when it assumes the role. You can use these tags to author a single role that can work across resources by allowing access to AWS resources based on matching tags. By default, EKS Pod Identity attaches six tags, including tags for cluster name, namespace, and service account name. For the list of tags added by EKS Pod Identity, see [List of session tags added by EKS Pod Identity](https://docs.aws.amazon.com/eks/latest/userguide/pod-id-abac.html#pod-id-abac-tags) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-podidentityassociation.html#cfn-eks-podidentityassociation-disablesessiontags
	//
	DisableSessionTags interface{} `field:"optional" json:"disableSessionTags" yaml:"disableSessionTags"`
	// Metadata that assists with categorization and organization.
	//
	// Each tag consists of a key and an optional value. You define both. Tags don't propagate to any other cluster or AWS resources.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource – 50
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length – 128 Unicode characters in UTF-8
	// - Maximum value length – 256 Unicode characters in UTF-8
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-podidentityassociation.html#cfn-eks-podidentityassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the target IAM role to associate with the service account.
	//
	// This role is assumed by using the EKS Pod Identity association role, then the credentials for this role are injected into the Pod.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-podidentityassociation.html#cfn-eks-podidentityassociation-targetrolearn
	//
	TargetRoleArn *string `field:"optional" json:"targetRoleArn" yaml:"targetRoleArn"`
}

