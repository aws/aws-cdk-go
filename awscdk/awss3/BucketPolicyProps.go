package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   bucketPolicyProps := &BucketPolicyProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   }
//
type BucketPolicyProps struct {
	// The Amazon S3 bucket that the policy applies to.
	Bucket IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Policy to apply when the policy is removed from this stack.
	// Default: - RemovalPolicy.DESTROY.
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

