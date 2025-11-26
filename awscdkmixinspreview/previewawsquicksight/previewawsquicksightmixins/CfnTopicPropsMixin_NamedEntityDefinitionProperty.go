package previewawsquicksightmixins


// A structure that represents a named entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   namedEntityDefinitionProperty := &NamedEntityDefinitionProperty{
//   	FieldName: jsii.String("fieldName"),
//   	Metric: &NamedEntityDefinitionMetricProperty{
//   		Aggregation: jsii.String("aggregation"),
//   		AggregationFunctionParameters: map[string]*string{
//   			"aggregationFunctionParametersKey": jsii.String("aggregationFunctionParameters"),
//   		},
//   	},
//   	PropertyName: jsii.String("propertyName"),
//   	PropertyRole: jsii.String("propertyRole"),
//   	PropertyUsage: jsii.String("propertyUsage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-namedentitydefinition.html
//
type CfnTopicPropsMixin_NamedEntityDefinitionProperty struct {
	// The name of the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-namedentitydefinition.html#cfn-quicksight-topic-namedentitydefinition-fieldname
	//
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// The definition of a metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-namedentitydefinition.html#cfn-quicksight-topic-namedentitydefinition-metric
	//
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// The property name to be used for the named entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-namedentitydefinition.html#cfn-quicksight-topic-namedentitydefinition-propertyname
	//
	PropertyName *string `field:"optional" json:"propertyName" yaml:"propertyName"`
	// The property role.
	//
	// Valid values for this structure are `PRIMARY` and `ID` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-namedentitydefinition.html#cfn-quicksight-topic-namedentitydefinition-propertyrole
	//
	PropertyRole *string `field:"optional" json:"propertyRole" yaml:"propertyRole"`
	// The property usage.
	//
	// Valid values for this structure are `INHERIT` , `DIMENSION` , and `MEASURE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-namedentitydefinition.html#cfn-quicksight-topic-namedentitydefinition-propertyusage
	//
	PropertyUsage *string `field:"optional" json:"propertyUsage" yaml:"propertyUsage"`
}

