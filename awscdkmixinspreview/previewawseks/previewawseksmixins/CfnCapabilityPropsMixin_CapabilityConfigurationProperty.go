package previewawseksmixins


// Configuration settings for a capability.
//
// The structure of this object varies depending on the capability type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capabilityConfigurationProperty := &CapabilityConfigurationProperty{
//   	ArgoCd: &ArgoCdProperty{
//   		AwsIdc: &AwsIdcProperty{
//   			IdcInstanceArn: jsii.String("idcInstanceArn"),
//   			IdcManagedApplicationArn: jsii.String("idcManagedApplicationArn"),
//   			IdcRegion: jsii.String("idcRegion"),
//   		},
//   		Namespace: jsii.String("namespace"),
//   		NetworkAccess: &NetworkAccessProperty{
//   			VpceIds: []*string{
//   				jsii.String("vpceIds"),
//   			},
//   		},
//   		RbacRoleMappings: []interface{}{
//   			&ArgoCdRoleMappingProperty{
//   				Identities: []interface{}{
//   					&SsoIdentityProperty{
//   						Id: jsii.String("id"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				Role: jsii.String("role"),
//   			},
//   		},
//   		ServerUrl: jsii.String("serverUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-capabilityconfiguration.html
//
type CfnCapabilityPropsMixin_CapabilityConfigurationProperty struct {
	// Configuration settings for an Argo CD capability.
	//
	// This includes the Kubernetes namespace, IAM Identity Center integration, RBAC role mappings, and network access configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-capabilityconfiguration.html#cfn-eks-capability-capabilityconfiguration-argocd
	//
	ArgoCd interface{} `field:"optional" json:"argoCd" yaml:"argoCd"`
}

