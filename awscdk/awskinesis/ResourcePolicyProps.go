package awskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties to associate a data stream with a policy.
//
// Example:
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//
//   // create a custom policy document
//   policyDocument := iam.NewPolicyDocument(&PolicyDocumentProps{
//   	AssignSids: jsii.Boolean(true),
//   	Statements: []policyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("kinesis:GetRecords"),
//   			},
//   			Resources: []*string{
//   				stream.StreamArn,
//   			},
//   			Principals: []iPrincipal{
//   				iam.NewAnyPrincipal(),
//   			},
//   		}),
//   	},
//   })
//
//   // create a resource policy manually
//   // create a resource policy manually
//   kinesis.NewResourcePolicy(this, jsii.String("ResourcePolicy"), &ResourcePolicyProps{
//   	Stream: Stream,
//   	PolicyDocument: PolicyDocument,
//   })
//
type ResourcePolicyProps struct {
	// The stream this policy applies to.
	Stream IStream `field:"required" json:"stream" yaml:"stream"`
	// IAM policy document to apply to a data stream.
	// Default: - empty policy document.
	//
	PolicyDocument awsiam.PolicyDocument `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

