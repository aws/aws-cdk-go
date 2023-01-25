package awsiotthingsgraph


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
	// `CfnFlowTemplate.DefinitionDocumentProperty.Language`.
	Language *string `field:"required" json:"language" yaml:"language"`
	// `CfnFlowTemplate.DefinitionDocumentProperty.Text`.
	Text *string `field:"required" json:"text" yaml:"text"`
}

