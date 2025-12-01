package awseks


// Configuration settings for an Argo CD capability.
//
// This includes the Kubernetes namespace, IAM Identity Center integration, RBAC role mappings, and network access configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   argoCdProperty := &ArgoCdProperty{
//   	AwsIdc: &AwsIdcProperty{
//   		IdcInstanceArn: jsii.String("idcInstanceArn"),
//
//   		// the properties below are optional
//   		IdcManagedApplicationArn: jsii.String("idcManagedApplicationArn"),
//   		IdcRegion: jsii.String("idcRegion"),
//   	},
//
//   	// the properties below are optional
//   	Namespace: jsii.String("namespace"),
//   	NetworkAccess: &NetworkAccessProperty{
//   		VpceIds: []*string{
//   			jsii.String("vpceIds"),
//   		},
//   	},
//   	RbacRoleMappings: []interface{}{
//   		&ArgoCdRoleMappingProperty{
//   			Identities: []interface{}{
//   				&SsoIdentityProperty{
//   					Id: jsii.String("id"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Role: jsii.String("role"),
//   		},
//   	},
//   	ServerUrl: jsii.String("serverUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-argocd.html
//
type CfnCapability_ArgoCdProperty struct {
	// Configuration for integrating Argo CD with IAM Identity Center.
	//
	// This allows you to use your organization's identity provider for authentication to Argo CD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-argocd.html#cfn-eks-capability-argocd-awsidc
	//
	AwsIdc interface{} `field:"required" json:"awsIdc" yaml:"awsIdc"`
	// The Kubernetes namespace where Argo CD resources will be created.
	//
	// If not specified, the default namespace is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-argocd.html#cfn-eks-capability-argocd-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Configuration for network access to the Argo CD capability's managed API server endpoint.
	//
	// By default, the Argo CD server is accessible via a public endpoint. You can optionally specify one or more VPC endpoint IDs to enable private connectivity from your VPCs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-argocd.html#cfn-eks-capability-argocd-networkaccess
	//
	NetworkAccess interface{} `field:"optional" json:"networkAccess" yaml:"networkAccess"`
	// A list of role mappings that define which IAM Identity Center users or groups have which Argo CD roles.
	//
	// Each mapping associates an Argo CD role (ADMIN, EDITOR, or VIEWER) with one or more IAM Identity Center identities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-argocd.html#cfn-eks-capability-argocd-rbacrolemappings
	//
	RbacRoleMappings interface{} `field:"optional" json:"rbacRoleMappings" yaml:"rbacRoleMappings"`
	// The URL of the Argo CD server.
	//
	// Use this URL to access the Argo CD web interface and API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-argocd.html#cfn-eks-capability-argocd-serverurl
	//
	ServerUrl *string `field:"optional" json:"serverUrl" yaml:"serverUrl"`
}

