package awsiotthingsgraph


// Properties for defining a `CfnFlowTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowTemplateProps := &cfnFlowTemplateProps{
//   	definition: &definitionDocumentProperty{
//   		language: jsii.String("language"),
//   		text: jsii.String("text"),
//   	},
//
//   	// the properties below are optional
//   	compatibleNamespaceVersion: jsii.Number(123),
//   }
//
type CfnFlowTemplateProps struct {
	// A workflow's definition document.
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// The version of the user's namespace against which the workflow was validated.
	//
	// Use this value in your system instance.
	CompatibleNamespaceVersion *float64 `field:"optional" json:"compatibleNamespaceVersion" yaml:"compatibleNamespaceVersion"`
}

