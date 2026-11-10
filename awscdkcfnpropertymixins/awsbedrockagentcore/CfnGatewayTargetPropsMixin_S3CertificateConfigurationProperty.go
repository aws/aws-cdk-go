package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3CertificateConfigurationProperty := &S3CertificateConfigurationProperty{
//   	BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-s3certificateconfiguration.html
//
type CfnGatewayTargetPropsMixin_S3CertificateConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-s3certificateconfiguration.html#cfn-bedrockagentcore-gatewaytarget-s3certificateconfiguration-bucketowneraccountid
	//
	BucketOwnerAccountId *string `field:"optional" json:"bucketOwnerAccountId" yaml:"bucketOwnerAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-s3certificateconfiguration.html#cfn-bedrockagentcore-gatewaytarget-s3certificateconfiguration-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

