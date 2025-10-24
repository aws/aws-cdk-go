package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The mTLS authentication configuration for a custom domain name.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//   var bucket Bucket
//
//
//   certArn := "arn:aws:acm:us-east-1:111111111111:certificate"
//   domainName := "example.com"
//
//   apigwv2.NewDomainName(this, jsii.String("DomainName"), &DomainNameProps{
//   	DomainName: jsii.String(DomainName),
//   	Certificate: acm.Certificate_FromCertificateArn(this, jsii.String("cert"), certArn),
//   	Mtls: &MTLSConfig{
//   		Bucket: *Bucket,
//   		Key: jsii.String("someca.pem"),
//   		Version: jsii.String("version"),
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
	// Default: - latest version.
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

