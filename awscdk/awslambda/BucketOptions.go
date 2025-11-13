package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Optional parameters for creating code using bucket.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   var key Key
//
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//
//   options := map[string]Key{
//   	"sourceKMSKey": key,
//   }
//   fnBucket := lambda.NewFunction(this, jsii.String("myFunction2"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromBucketV2(bucket, jsii.String("python-lambda-handler.zip"), options),
//   })
//
type BucketOptions struct {
	// Optional S3 object version.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
	// The ARN of the KMS key used to encrypt the handler code.
	// Default: - the default server-side encryption with Amazon S3 managed keys(SSE-S3) key will be used.
	//
	SourceKMSKey interfacesawskms.IKeyRef `field:"optional" json:"sourceKMSKey" yaml:"sourceKMSKey"`
}

