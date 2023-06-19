package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Result of calling addToResourcePolicy.
//
// Example:
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("existingBucket"), jsii.String("bucket-name"))
//
//   // No policy statement will be added to the resource
//   result := bucket.AddToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("s3:GetObject"),
//   	},
//   	Resources: []*string{
//   		bucket.ArnForObjects(jsii.String("file.txt")),
//   	},
//   	Principals: []iPrincipal{
//   		iam.NewAccountRootPrincipal(),
//   	},
//   }))
//
// Experimental.
type AddToResourcePolicyResult struct {
	// Whether the statement was added.
	// Experimental.
	StatementAdded *bool `field:"required" json:"statementAdded" yaml:"statementAdded"`
	// Dependable which allows depending on the policy change being applied.
	// Experimental.
	PolicyDependable awscdk.IDependable `field:"optional" json:"policyDependable" yaml:"policyDependable"`
}

