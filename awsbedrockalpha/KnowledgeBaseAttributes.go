package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

// Attributes for referencing an existing knowledge base.
//
// Example:
//   var role Role
//
//
//   byArn := bedrock.KnowledgeBase_FromKnowledgeBaseArn(this, jsii.String("ImportedByArn"), jsii.String("arn:aws:bedrock:us-east-1:123456789012:knowledge-base/KB12345678"))
//
//   // Assumed to be in the same account and region as this stack
//   byId := bedrock.KnowledgeBase_FromKnowledgeBaseId(this, jsii.String("ImportedById"), jsii.String("KB12345678"))
//
//   withRole := bedrock.KnowledgeBase_FromKnowledgeBaseAttributes(this, jsii.String("ImportedWithRole"), &KnowledgeBaseAttributes{
//   	KnowledgeBaseArn: jsii.String("arn:aws:bedrock:us-east-1:123456789012:knowledge-base/KB12345678"),
//   	Role: Role,
//   })
//
// Experimental.
type KnowledgeBaseAttributes struct {
	// The ARN of the knowledge base.
	// Experimental.
	KnowledgeBaseArn *string `field:"required" json:"knowledgeBaseArn" yaml:"knowledgeBaseArn"`
	// The service role of the knowledge base.
	// Default: - the role is unknown.
	//
	// Experimental.
	Role interface{ interfacesawsiam.IRoleRef;awsiam.IGrantable } `field:"optional" json:"role" yaml:"role"`
}

