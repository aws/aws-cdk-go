package awsekslegacy

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   	preserveOnDelete: jsii.Boolean(false),
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
	// The configuration values that you provided.
	ConfigurationValues *string `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	// `AWS::EKS::Addon.PreserveOnDelete`.
	PreserveOnDelete interface{} `field:"optional" json:"preserveOnDelete" yaml:"preserveOnDelete"`
	// How to resolve field value conflicts for an Amazon EKS add-on.
	//
	// Conflicts are handled based on the value you choose:
	//
	// - *None* – If the self-managed version of the add-on is installed on your cluster, Amazon EKS doesn't change the value. Creation of the add-on might fail.
	// - *Overwrite* – If the self-managed version of the add-on is installed on your cluster and the Amazon EKS default value is different than the existing value, Amazon EKS changes the value to the Amazon EKS default value.
	// - *Preserve* – Not supported. You can set this value when updating an add-on though. For more information, see [UpdateAddon](https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html) .
	//
	// If you don't currently have the self-managed version of the add-on installed on your cluster, the Amazon EKS add-on is installed. Amazon EKS sets all values to default values, regardless of the option that you specify.
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

