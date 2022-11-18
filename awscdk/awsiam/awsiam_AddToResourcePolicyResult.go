package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Result of calling addToResourcePolicy.
//
// Example:
//   bucket := s3.bucket.fromBucketName(this, jsii.String("existingBucket"), jsii.String("bucket-name"))
//
//   // No policy statement will be added to the resource
//   result := bucket.addToResourcePolicy(iam.NewPolicyStatement(&policyStatementProps{
//   	actions: []*string{
//   		jsii.String("s3:GetObject"),
//   	},
//   	resources: []*string{
//   		bucket.arnForObjects(jsii.String("file.txt")),
//   	},
//   	principals: []iPrincipal{
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

