package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Information about a revocation file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucketRef iBucketRef
//
//   revocationContent := &RevocationContent{
//   	Bucket: bucketRef,
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	RevocationType: awscdk.Aws_elasticloadbalancingv2.RevocationType_CRL,
//   	Version: jsii.String("version"),
//   }
//
type RevocationContent struct {
	// The Amazon S3 bucket for the revocation file.
	Bucket awss3.IBucketRef `field:"required" json:"bucket" yaml:"bucket"`
	// The Amazon S3 path for the revocation file.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The type of revocation file.
	// Default: RevocationType.CRL
	//
	RevocationType RevocationType `field:"optional" json:"revocationType" yaml:"revocationType"`
	// The Amazon S3 object version of the revocation file.
	// Default: - latest version.
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

