package previewawscloudfrontmixins


// Amazon CloudFront supports gRPC, an open-source remote procedure call (RPC) framework built on HTTP/2.
//
// gRPC offers bi-directional streaming and binary protocol that buffers payloads, making it suitable for applications that require low latency communications.
//
// To enable your distribution to handle gRPC requests, you must include HTTP/2 as one of the supported `HTTP` versions and allow `HTTP` methods, including `POST` .
//
// For more information, see [Using gRPC with CloudFront distributions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-using-grpc.html) in the *Amazon CloudFront Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   grpcConfigProperty := &GrpcConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-grpcconfig.html
//
type CfnDistributionPropsMixin_GrpcConfigProperty struct {
	// Enables your CloudFront distribution to receive gRPC requests and to proxy them directly to your origins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-grpcconfig.html#cfn-cloudfront-distribution-grpcconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

