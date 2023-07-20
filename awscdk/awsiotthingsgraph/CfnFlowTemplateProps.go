package awsiotthingsgraph


// Properties for defining a `CfnFlowTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowTemplateProps := &CfnFlowTemplateProps{
//   	Definition: &DefinitionDocumentProperty{
//   		Language: jsii.String("language"),
//   		Text: jsii.String("text"),
//   	},
//
//   	// the properties below are optional
//   	CompatibleNamespaceVersion: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html
//
type CfnFlowTemplateProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html#cfn-iotthingsgraph-flowtemplate-definition
	//
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html#cfn-iotthingsgraph-flowtemplate-compatiblenamespaceversion
	//
	CompatibleNamespaceVersion *float64 `field:"optional" json:"compatibleNamespaceVersion" yaml:"compatibleNamespaceVersion"`
}

