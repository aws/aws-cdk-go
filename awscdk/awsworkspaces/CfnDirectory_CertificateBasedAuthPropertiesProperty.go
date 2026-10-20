package awsworkspaces


// Describes the properties of the certificate-based authentication you want to use with your WorkSpaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateBasedAuthPropertiesProperty := &CertificateBasedAuthPropertiesProperty{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-certificatebasedauthproperties.html
//
type CfnDirectory_CertificateBasedAuthPropertiesProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Web Services Certificate Manager Private CA resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-certificatebasedauthproperties.html#cfn-workspaces-directory-certificatebasedauthproperties-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The status of the certificate-based authentication properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-certificatebasedauthproperties.html#cfn-workspaces-directory-certificatebasedauthproperties-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

