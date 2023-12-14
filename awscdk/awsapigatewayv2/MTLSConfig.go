package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The mTLS authentication configuration for a custom domain name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   mTLSConfig := &MTLSConfig{
//   	Bucket: bucket,
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	Version: jsii.String("version"),
//   }
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

