package awsquicksight


// A structure that represents a semantic type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   semanticTypeProperty := &SemanticTypeProperty{
//   	FalseyCellValue: jsii.String("falseyCellValue"),
//   	FalseyCellValueSynonyms: []*string{
//   		jsii.String("falseyCellValueSynonyms"),
//   	},
//   	SubTypeName: jsii.String("subTypeName"),
//   	TruthyCellValue: jsii.String("truthyCellValue"),
//   	TruthyCellValueSynonyms: []*string{
//   		jsii.String("truthyCellValueSynonyms"),
//   	},
//   	TypeName: jsii.String("typeName"),
//   	TypeParameters: map[string]*string{
//   		"typeParametersKey": jsii.String("typeParameters"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semantictype.html
//
type CfnTopic_SemanticTypeProperty struct {
	// The semantic type falsey cell value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semantictype.html#cfn-quicksight-topic-semantictype-falseycellvalue
	//
	FalseyCellValue *string `field:"optional" json:"falseyCellValue" yaml:"falseyCellValue"`
	// The other names or aliases for the false cell value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semantictype.html#cfn-quicksight-topic-semantictype-falseycellvaluesynonyms
	//
	FalseyCellValueSynonyms *[]*string `field:"optional" json:"falseyCellValueSynonyms" yaml:"falseyCellValueSynonyms"`
	// The semantic type sub type name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semantictype.html#cfn-quicksight-topic-semantictype-subtypename
	//
	SubTypeName *string `field:"optional" json:"subTypeName" yaml:"subTypeName"`
	// The semantic type truthy cell value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semantictype.html#cfn-quicksight-topic-semantictype-truthycellvalue
	//
	TruthyCellValue *string `field:"optional" json:"truthyCellValue" yaml:"truthyCellValue"`
	// The other names or aliases for the true cell value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semantictype.html#cfn-quicksight-topic-semantictype-truthycellvaluesynonyms
	//
	TruthyCellValueSynonyms *[]*string `field:"optional" json:"truthyCellValueSynonyms" yaml:"truthyCellValueSynonyms"`
	// The semantic type name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semantictype.html#cfn-quicksight-topic-semantictype-typename
	//
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// The semantic type parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semantictype.html#cfn-quicksight-topic-semantictype-typeparameters
	//
	TypeParameters interface{} `field:"optional" json:"typeParameters" yaml:"typeParameters"`
}

