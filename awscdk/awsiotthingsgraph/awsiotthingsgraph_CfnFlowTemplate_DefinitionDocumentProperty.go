package awsiotthingsgraph


// A document that defines an entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   definitionDocumentProperty := &definitionDocumentProperty{
//   	language: jsii.String("language"),
//   	text: jsii.String("text"),
//   }
//
type CfnFlowTemplate_DefinitionDocumentProperty struct {
	// The language used to define the entity.
	//
	// `GRAPHQL` is the only valid value.
	Language *string `field:"required" json:"language" yaml:"language"`
	// The GraphQL text that defines the entity.
	Text *string `field:"required" json:"text" yaml:"text"`
}

