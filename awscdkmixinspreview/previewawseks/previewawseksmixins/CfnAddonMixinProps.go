package previewawseksmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAddonPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAddonMixinProps := &CfnAddonMixinProps{
//   	AddonName: jsii.String("addonName"),
//   	AddonVersion: jsii.String("addonVersion"),
//   	ClusterName: jsii.String("clusterName"),
//   	ConfigurationValues: jsii.String("configurationValues"),
//   	NamespaceConfig: &NamespaceConfigProperty{
//   		Namespace: jsii.String("namespace"),
//   	},
//   	PodIdentityAssociations: []interface{}{
//   		&PodIdentityAssociationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			ServiceAccount: jsii.String("serviceAccount"),
//   		},
//   	},
//   	PreserveOnDelete: jsii.Boolean(false),
//   	ResolveConflicts: jsii.String("resolveConflicts"),
//   	ServiceAccountRoleArn: jsii.String("serviceAccountRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html
//
type CfnAddonMixinProps struct {
	// The name of the add-on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-addonname
	//
	AddonName *string `field:"optional" json:"addonName" yaml:"addonName"`
	// The version of the add-on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-addonversion
	//
	AddonVersion *string `field:"optional" json:"addonVersion" yaml:"addonVersion"`
	// The name of your cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The configuration values that you provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-configurationvalues
	//
	ConfigurationValues *string `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	// The namespace configuration for the addon.
	//
	// This specifies the Kubernetes namespace where the addon is installed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-namespaceconfig
	//
	NamespaceConfig interface{} `field:"optional" json:"namespaceConfig" yaml:"namespaceConfig"`
	// An array of EKS Pod Identity associations owned by the add-on.
	//
	// Each association maps a role to a service account in a namespace in the cluster.
	//
	// For more information, see [Attach an IAM Role to an Amazon EKS add-on using EKS Pod Identity](https://docs.aws.amazon.com/eks/latest/userguide/add-ons-iam.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-podidentityassociations
	//
	PodIdentityAssociations interface{} `field:"optional" json:"podIdentityAssociations" yaml:"podIdentityAssociations"`
	// Specifying this option preserves the add-on software on your cluster but Amazon EKS stops managing any settings for the add-on.
	//
	// If an IAM account is associated with the add-on, it isn't removed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-preserveondelete
	//
	PreserveOnDelete interface{} `field:"optional" json:"preserveOnDelete" yaml:"preserveOnDelete"`
	// How to resolve field value conflicts for an Amazon EKS add-on.
	//
	// Conflicts are handled based on the value you choose:
	//
	// - *None* – If the self-managed version of the add-on is installed on your cluster, Amazon EKS doesn't change the value. Creation of the add-on might fail.
	// - *Overwrite* – If the self-managed version of the add-on is installed on your cluster and the Amazon EKS default value is different than the existing value, Amazon EKS changes the value to the Amazon EKS default value.
	// - *Preserve* – This is similar to the NONE option. If the self-managed version of the add-on is installed on your cluster Amazon EKS doesn't change the add-on resource properties. Creation of the add-on might fail if conflicts are detected. This option works differently during the update operation. For more information, see [`UpdateAddon`](https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html) .
	//
	// If you don't currently have the self-managed version of the add-on installed on your cluster, the Amazon EKS add-on is installed. Amazon EKS sets all values to default values, regardless of the option that you specify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-resolveconflicts
	//
	ResolveConflicts *string `field:"optional" json:"resolveConflicts" yaml:"resolveConflicts"`
	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.
	//
	// The role must be assigned the IAM permissions required by the add-on. If you don't specify an existing IAM role, then the add-on uses the permissions assigned to the node IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the *Amazon EKS User Guide* .
	//
	// > To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC) provider created for your cluster. For more information, see [Enabling IAM roles for service accounts on your cluster](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html) in the *Amazon EKS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-serviceaccountrolearn
	//
	ServiceAccountRoleArn *string `field:"optional" json:"serviceAccountRoleArn" yaml:"serviceAccountRoleArn"`
	// The metadata that you apply to the add-on to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Add-on tags do not propagate to any other resources associated with the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

