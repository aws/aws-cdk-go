package previewawsec2mixins


// Properties for CfnEnclaveCertificateIamRoleAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnclaveCertificateIamRoleAssociationMixinProps := &CfnEnclaveCertificateIamRoleAssociationMixinProps{
//   	CertificateArn: jsii.String("certificateArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-enclavecertificateiamroleassociation.html
//
type CfnEnclaveCertificateIamRoleAssociationMixinProps struct {
	// The ARN of the ACM certificate with which to associate the IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-enclavecertificateiamroleassociation.html#cfn-ec2-enclavecertificateiamroleassociation-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The ARN of the IAM role to associate with the ACM certificate.
	//
	// You can associate up to 16 IAM roles with an ACM certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-enclavecertificateiamroleassociation.html#cfn-ec2-enclavecertificateiamroleassociation-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

