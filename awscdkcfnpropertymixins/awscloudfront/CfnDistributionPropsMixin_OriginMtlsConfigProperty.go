package awscloudfront


// Configures mutual TLS authentication between CloudFront and your origin server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   originMtlsConfigProperty := &OriginMtlsConfigProperty{
//   	ClientCertificateArn: jsii.String("clientCertificateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-originmtlsconfig.html
//
type CfnDistributionPropsMixin_OriginMtlsConfigProperty struct {
	// The Amazon Resource Name (ARN) of the client certificate stored in AWS Certificate Manager (ACM) that CloudFront uses to authenticate with your origin using Mutual TLS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-originmtlsconfig.html#cfn-cloudfront-distribution-originmtlsconfig-clientcertificatearn
	//
	ClientCertificateArn *string `field:"optional" json:"clientCertificateArn" yaml:"clientCertificateArn"`
}

