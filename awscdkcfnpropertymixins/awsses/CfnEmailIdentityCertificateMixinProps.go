package awsses


// Properties for CfnEmailIdentityCertificatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnEmailIdentityCertificateMixinProps := &CfnEmailIdentityCertificateMixinProps{
//   	CertificateArn: jsii.String("certificateArn"),
//   	EmailIdentity: jsii.String("emailIdentity"),
//   	FromAddress: jsii.String("fromAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentitycertificate.html
//
type CfnEmailIdentityCertificateMixinProps struct {
	// The ARN of the AWS Certificate Manager certificate to associate with the sender.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentitycertificate.html#cfn-ses-emailidentitycertificate-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The email identity that owns the sender the certificate is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentitycertificate.html#cfn-ses-emailidentitycertificate-emailidentity
	//
	EmailIdentity *string `field:"optional" json:"emailIdentity" yaml:"emailIdentity"`
	// The sender the certificate signs for.
	//
	// For an email address identity this is the identity itself.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentitycertificate.html#cfn-ses-emailidentitycertificate-fromaddress
	//
	FromAddress *string `field:"optional" json:"fromAddress" yaml:"fromAddress"`
}

