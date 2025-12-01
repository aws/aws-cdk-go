package previewawseksmixins


// A mapping between an Argo CD role and IAM Identity Center identities.
//
// This defines which users or groups have specific permissions in Argo CD.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   argoCdRoleMappingProperty := &ArgoCdRoleMappingProperty{
//   	Identities: []interface{}{
//   		&SsoIdentityProperty{
//   			Id: jsii.String("id"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Role: jsii.String("role"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-argocdrolemapping.html
//
type CfnCapabilityPropsMixin_ArgoCdRoleMappingProperty struct {
	// A list of IAM Identity Center identities (users or groups) that should be assigned this Argo CD role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-argocdrolemapping.html#cfn-eks-capability-argocdrolemapping-identities
	//
	Identities interface{} `field:"optional" json:"identities" yaml:"identities"`
	// The Argo CD role to assign.
	//
	// Valid values are: ADMIN (full administrative access to Argo CD), EDITOR (edit access to Argo CD resources), or VIEWER (read-only access to Argo CD resources).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-argocdrolemapping.html#cfn-eks-capability-argocdrolemapping-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
}

