package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3CertificateConfigurationProperty := &S3CertificateConfigurationProperty{
//   	Uri: jsii.String("uri"),
//
//   	// the properties below are optional
//   	BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-s3certificateconfiguration.html
//
type CfnGatewayTarget_S3CertificateConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-s3certificateconfiguration.html#cfn-bedrockagentcore-gatewaytarget-s3certificateconfiguration-uri
	//
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-s3certificateconfiguration.html#cfn-bedrockagentcore-gatewaytarget-s3certificateconfiguration-bucketowneraccountid
	//
	BucketOwnerAccountId *string `field:"optional" json:"bucketOwnerAccountId" yaml:"bucketOwnerAccountId"`
}

