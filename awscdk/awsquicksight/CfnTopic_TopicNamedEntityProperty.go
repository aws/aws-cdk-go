package awsquicksight


// A structure that represents a named entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicNamedEntityProperty := &TopicNamedEntityProperty{
//   	EntityName: jsii.String("entityName"),
//
//   	// the properties below are optional
//   	Definition: []interface{}{
//   		&NamedEntityDefinitionProperty{
//   			FieldName: jsii.String("fieldName"),
//   			Metric: &NamedEntityDefinitionMetricProperty{
//   				Aggregation: jsii.String("aggregation"),
//   				AggregationFunctionParameters: map[string]*string{
//   					"aggregationFunctionParametersKey": jsii.String("aggregationFunctionParameters"),
//   				},
//   			},
//   			PropertyName: jsii.String("propertyName"),
//   			PropertyRole: jsii.String("propertyRole"),
//   			PropertyUsage: jsii.String("propertyUsage"),
//   		},
//   	},
//   	EntityDescription: jsii.String("entityDescription"),
//   	EntitySynonyms: []*string{
//   		jsii.String("entitySynonyms"),
//   	},
//   	SemanticEntityType: &SemanticEntityTypeProperty{
//   		SubTypeName: jsii.String("subTypeName"),
//   		TypeName: jsii.String("typeName"),
//   		TypeParameters: map[string]*string{
//   			"typeParametersKey": jsii.String("typeParameters"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnamedentity.html
//
type CfnTopic_TopicNamedEntityProperty struct {
	// The name of the named entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnamedentity.html#cfn-quicksight-topic-topicnamedentity-entityname
	//
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// The definition of a named entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnamedentity.html#cfn-quicksight-topic-topicnamedentity-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// The description of the named entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnamedentity.html#cfn-quicksight-topic-topicnamedentity-entitydescription
	//
	EntityDescription *string `field:"optional" json:"entityDescription" yaml:"entityDescription"`
	// The other names or aliases for the named entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnamedentity.html#cfn-quicksight-topic-topicnamedentity-entitysynonyms
	//
	EntitySynonyms *[]*string `field:"optional" json:"entitySynonyms" yaml:"entitySynonyms"`
	// The type of named entity that a topic represents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnamedentity.html#cfn-quicksight-topic-topicnamedentity-semanticentitytype
	//
	SemanticEntityType interface{} `field:"optional" json:"semanticEntityType" yaml:"semanticEntityType"`
}

