package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originMtlsConfigProperty := &OriginMtlsConfigProperty{
//   	ClientCertificateArn: jsii.String("clientCertificateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-originmtlsconfig.html
//
type CfnDistribution_OriginMtlsConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-originmtlsconfig.html#cfn-cloudfront-distribution-originmtlsconfig-clientcertificatearn
	//
	ClientCertificateArn *string `field:"required" json:"clientCertificateArn" yaml:"clientCertificateArn"`
}

