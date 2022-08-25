package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The mTLS authentication configuration for a custom domain name.
//
// Example:
//   var acm interface{}
//
//
//   apigateway.NewDomainName(this, jsii.String("domain-name"), &domainNameProps{
//   	domainName: jsii.String("example.com"),
//   	certificate: acm.certificate_FromCertificateArn(this, jsii.String("cert"), jsii.String("arn:aws:acm:us-east-1:1111111:certificate/11-3336f1-44483d-adc7-9cd375c5169d")),
//   	mtls: &mTLSConfig{
//   		bucket: s3.NewBucket(this, jsii.String("bucket")),
//   		key: jsii.String("truststore.pem"),
//   		version: jsii.String("version"),
//   	},
//   })
//
type MTLSConfig struct {
	// The bucket that the trust store is hosted in.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The key in S3 to look at for the trust store.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The version of the S3 object that contains your truststore.
	//
	// To specify a version, you must have versioning enabled for the S3 bucket.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

