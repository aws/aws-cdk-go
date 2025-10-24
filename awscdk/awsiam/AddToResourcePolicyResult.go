package awsiam

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Result of calling addToResourcePolicy.
//
// Example:
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("existingBucket"), jsii.String("amzn-s3-demo-bucket"))
//
//   // No policy statement will be added to the resource
//   result := bucket.AddToResourcePolicy(
//   iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("s3:GetObject"),
//   	},
//   	Resources: []*string{
//   		bucket.ArnForObjects(jsii.String("file.txt")),
//   	},
//   	Principals: []IPrincipal{
//   		iam.NewAccountRootPrincipal(),
//   	},
//   }))
//
type AddToResourcePolicyResult struct {
	// Whether the statement was added.
	StatementAdded *bool `field:"required" json:"statementAdded" yaml:"statementAdded"`
	// Dependable which allows depending on the policy change being applied.
	// Default: - If `statementAdded` is true, the resource object itself.
	// Otherwise, no dependable.
	//
	PolicyDependable constructs.IDependable `field:"optional" json:"policyDependable" yaml:"policyDependable"`
}

