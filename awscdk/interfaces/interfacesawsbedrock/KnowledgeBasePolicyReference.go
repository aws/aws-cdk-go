package interfacesawsbedrock


// A reference to a KnowledgeBasePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   knowledgeBasePolicyReference := &KnowledgeBasePolicyReference{
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   }
//
type KnowledgeBasePolicyReference struct {
	// The KnowledgeBaseId of the KnowledgeBasePolicy resource.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
}

