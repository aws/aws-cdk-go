package awsappstream


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateBasedAuthPropertiesProperty := &certificateBasedAuthPropertiesProperty{
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	status: jsii.String("status"),
//   }
//
type CfnDirectoryConfig_CertificateBasedAuthPropertiesProperty struct {
	// `CfnDirectoryConfig.CertificateBasedAuthPropertiesProperty.CertificateAuthorityArn`.
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// `CfnDirectoryConfig.CertificateBasedAuthPropertiesProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

