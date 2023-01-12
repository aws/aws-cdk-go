package awsiam

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Result of calling addToResourcePolicy.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type AddToResourcePolicyResult struct {
	// Whether the statement was added.
	StatementAdded *bool `field:"required" json:"statementAdded" yaml:"statementAdded"`
	// Dependable which allows depending on the policy change being applied.
	PolicyDependable constructs.IDependable `field:"optional" json:"policyDependable" yaml:"policyDependable"`
}

