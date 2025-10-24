package awscloudfront


// VPC origin endpoint configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOriginOptions := &VpcOriginOptions{
//   	HttpPort: jsii.Number(123),
//   	HttpsPort: jsii.Number(123),
//   	OriginSslProtocols: []OriginSslPolicy{
//   		awscdk.Aws_cloudfront.OriginSslPolicy_SSL_V3,
//   	},
//   	ProtocolPolicy: awscdk.*Aws_cloudfront.OriginProtocolPolicy_HTTP_ONLY,
//   	VpcOriginName: jsii.String("vpcOriginName"),
//   }
//
type VpcOriginOptions struct {
	// The HTTP port for the CloudFront VPC origin endpoint configuration.
	// Default: 80.
	//
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The HTTPS port of the CloudFront VPC origin endpoint configuration.
	// Default: 443.
	//
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// A list that contains allowed SSL/TLS protocols for this distribution.
	// Default: - TLSv1.2
	//
	OriginSslProtocols *[]OriginSslPolicy `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
	// The origin protocol policy for the CloudFront VPC origin endpoint configuration.
	// Default: OriginProtocolPolicy.MATCH_VIEWER
	//
	ProtocolPolicy OriginProtocolPolicy `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
	// The name of the CloudFront VPC origin endpoint configuration.
	// Default: - generated from the `id`.
	//
	VpcOriginName *string `field:"optional" json:"vpcOriginName" yaml:"vpcOriginName"`
}

