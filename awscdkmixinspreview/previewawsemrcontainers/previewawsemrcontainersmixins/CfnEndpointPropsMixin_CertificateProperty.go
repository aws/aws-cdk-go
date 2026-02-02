package previewawsemrcontainersmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   certificateProperty := &CertificateProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	CertificateData: jsii.String("certificateData"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-certificate.html
//
type CfnEndpointPropsMixin_CertificateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-certificate.html#cfn-emrcontainers-endpoint-certificate-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-certificate.html#cfn-emrcontainers-endpoint-certificate-certificatedata
	//
	CertificateData *string `field:"optional" json:"certificateData" yaml:"certificateData"`
}

