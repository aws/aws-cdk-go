package interfacesawsbedrock


// A reference to a KnowledgeBase resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   knowledgeBaseReference := &KnowledgeBaseReference{
//   	KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   }
//
type KnowledgeBaseReference struct {
	// The ARN of the KnowledgeBase resource.
	KnowledgeBaseArn *string `field:"required" json:"knowledgeBaseArn" yaml:"knowledgeBaseArn"`
	// The KnowledgeBaseId of the KnowledgeBase resource.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
}

