package awskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties to associate a data stream with a policy.
//
// Example:
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//   streamConsumer := kinesis.NewStreamConsumer(this, jsii.String("MyStreamConsumer"), &StreamConsumerProps{
//   	StreamConsumerName: jsii.String("MyStreamConsumer"),
//   	Stream: Stream,
//   })
//
//   // create a custom policy document
//   policyDocument := iam.NewPolicyDocument(&PolicyDocumentProps{
//   	AssignSids: jsii.Boolean(true),
//   	Statements: []PolicyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("kinesis:GetRecords"),
//   			},
//   			Resources: []*string{
//   				stream.streamArn,
//   			},
//   			Principals: []IPrincipal{
//   				iam.NewAnyPrincipal(),
//   			},
//   		}),
//   	},
//   })
//
//   // create a stream resource policy manually
//   // create a stream resource policy manually
//   kinesis.NewResourcePolicy(this, jsii.String("ResourcePolicy"), &ResourcePolicyProps{
//   	Stream: Stream,
//   	PolicyDocument: PolicyDocument,
//   })
//
//   // create a stream consumer resource policy manually
//   // create a stream consumer resource policy manually
//   kinesis.NewResourcePolicy(this, jsii.String("ResourcePolicy"), &ResourcePolicyProps{
//   	StreamConsumer: StreamConsumer,
//   	PolicyDocument: PolicyDocument,
//   })
//
type ResourcePolicyProps struct {
	// IAM policy document to apply to a data stream.
	// Default: - empty policy document.
	//
	PolicyDocument awsiam.PolicyDocument `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The stream this policy applies to.
	//
	// Note: only one of `stream` and `streamConsumer` must be set.
	// Default: - policy is not associated to a stream.
	//
	Stream IStream `field:"optional" json:"stream" yaml:"stream"`
	// The stream consumer this policy applies to.
	//
	// Note: only one of `stream` and `streamConsumer` must be set.
	// Default: - policy is not associated to a consumer.
	//
	StreamConsumer IStreamConsumer `field:"optional" json:"streamConsumer" yaml:"streamConsumer"`
}

