package mixinsawsiotthingsgraph


// Properties for CfnFlowTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowTemplateMixinProps := &CfnFlowTemplateMixinProps{
//   	CompatibleNamespaceVersion: jsii.Number(123),
//   	Definition: &DefinitionDocumentProperty{
//   		Language: jsii.String("language"),
//   		Text: jsii.String("text"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html
//
type CfnFlowTemplateMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html#cfn-iotthingsgraph-flowtemplate-compatiblenamespaceversion
	//
	CompatibleNamespaceVersion *float64 `field:"optional" json:"compatibleNamespaceVersion" yaml:"compatibleNamespaceVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html#cfn-iotthingsgraph-flowtemplate-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
}

