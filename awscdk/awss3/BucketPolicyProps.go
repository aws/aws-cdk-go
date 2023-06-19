package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   bucketPolicyProps := &BucketPolicyProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	RemovalPolicy: monocdk.RemovalPolicy_DESTROY,
//   }
//
// Experimental.
type BucketPolicyProps struct {
	// The Amazon S3 bucket that the policy applies to.
	// Experimental.
	Bucket IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Policy to apply when the policy is removed from this stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

