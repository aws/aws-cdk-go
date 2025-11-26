package previewawsquicksightmixins


// A structure that represents a semantic entity type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   semanticEntityTypeProperty := &SemanticEntityTypeProperty{
//   	SubTypeName: jsii.String("subTypeName"),
//   	TypeName: jsii.String("typeName"),
//   	TypeParameters: map[string]*string{
//   		"typeParametersKey": jsii.String("typeParameters"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semanticentitytype.html
//
type CfnTopicPropsMixin_SemanticEntityTypeProperty struct {
	// The semantic entity sub type name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semanticentitytype.html#cfn-quicksight-topic-semanticentitytype-subtypename
	//
	SubTypeName *string `field:"optional" json:"subTypeName" yaml:"subTypeName"`
	// The semantic entity type name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semanticentitytype.html#cfn-quicksight-topic-semanticentitytype-typename
	//
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// The semantic entity type parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semanticentitytype.html#cfn-quicksight-topic-semanticentitytype-typeparameters
	//
	TypeParameters interface{} `field:"optional" json:"typeParameters" yaml:"typeParameters"`
}

