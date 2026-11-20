package awsbedrockagentcore


// A reference to a PEM-encoded private CA certificate stored as an Amazon S3 object.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-s3certificateconfiguration.html
//
type CfnOAuth2CredentialProviderPropsMixin_S3CertificateConfigurationProperty struct {
	// The account ID of the Amazon S3 bucket owner, used for cross-account access to the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-s3certificateconfiguration.html#cfn-bedrockagentcore-oauth2credentialprovider-s3certificateconfiguration-bucketowneraccountid
	//
	BucketOwnerAccountId *string `field:"optional" json:"bucketOwnerAccountId" yaml:"bucketOwnerAccountId"`
	// The URI of the Amazon S3 object that contains the PEM-encoded certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-s3certificateconfiguration.html#cfn-bedrockagentcore-oauth2credentialprovider-s3certificateconfiguration-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

