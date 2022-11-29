package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// The mTLS authentication configuration for a custom domain name.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//   var bucket bucket
//
//
//   certArn := "arn:aws:acm:us-east-1:111111111111:certificate"
//   domainName := "example.com"
//
//   apigwv2.NewDomainName(this, jsii.String("DomainName"), &domainNameProps{
//   	domainName: jsii.String(domainName),
//   	certificate: acm.certificate.fromCertificateArn(this, jsii.String("cert"), certArn),
//   	mtls: &mTLSConfig{
//   		bucket: bucket,
//   		key: jsii.String("someca.pem"),
//   		version: jsii.String("version"),
//   	},
//   })
//
// Experimental.
type MTLSConfig struct {
	// The bucket that the trust store is hosted in.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The key in S3 to look at for the trust store.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The version of the S3 object that contains your truststore.
	//
	// To specify a version, you must have versioning enabled for the S3 bucket.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

