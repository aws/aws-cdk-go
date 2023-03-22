package awsec2


// Properties for defining a `CfnEnclaveCertificateIamRoleAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnclaveCertificateIamRoleAssociationProps := &cfnEnclaveCertificateIamRoleAssociationProps{
//   	certificateArn: jsii.String("certificateArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnEnclaveCertificateIamRoleAssociationProps struct {
	// The ARN of the ACM certificate with which to associate the IAM role.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// The ARN of the IAM role to associate with the ACM certificate.
	//
	// You can associate up to 16 IAM roles with an ACM certificate.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

