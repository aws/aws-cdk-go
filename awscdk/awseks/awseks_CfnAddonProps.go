package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAddon`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAddonProps := &cfnAddonProps{
//   	addonName: jsii.String("addonName"),
//   	clusterName: jsii.String("clusterName"),
//
//   	// the properties below are optional
//   	addonVersion: jsii.String("addonVersion"),
//   	configurationValues: jsii.String("configurationValues"),
//   	resolveConflicts: jsii.String("resolveConflicts"),
//   	serviceAccountRoleArn: jsii.String("serviceAccountRoleArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAddonProps struct {
	// The name of the add-on.
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// The name of the cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The version of the add-on.
	AddonVersion *string `field:"optional" json:"addonVersion" yaml:"addonVersion"`
	// `AWS::EKS::Addon.ConfigurationValues`.
	ConfigurationValues *string `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	// How to resolve parameter value conflicts when migrating an existing add-on to an Amazon EKS add-on.
	ResolveConflicts *string `field:"optional" json:"resolveConflicts" yaml:"resolveConflicts"`
	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.
	//
	// The role must be assigned the IAM permissions required by the add-on. If you don't specify an existing IAM role, then the add-on uses the permissions assigned to the node IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the *Amazon EKS User Guide* .
	//
	// > To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC) provider created for your cluster. For more information, see [Enabling IAM roles for service accounts on your cluster](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html) in the *Amazon EKS User Guide* .
	ServiceAccountRoleArn *string `field:"optional" json:"serviceAccountRoleArn" yaml:"serviceAccountRoleArn"`
	// The metadata that you apply to the add-on to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Add-on tags do not propagate to any other resources associated with the cluster.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

