package awsqbusiness


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inlineDocumentEnrichmentConfigurationProperty := &InlineDocumentEnrichmentConfigurationProperty{
//   	Condition: &DocumentAttributeConditionProperty{
//   		Key: jsii.String("key"),
//   		Operator: jsii.String("operator"),
//
//   		// the properties below are optional
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	DocumentContentOperator: jsii.String("documentContentOperator"),
//   	Target: &DocumentAttributeTargetProperty{
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		AttributeValueOperator: jsii.String("attributeValueOperator"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html
//
type CfnDataSource_InlineDocumentEnrichmentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html#cfn-qbusiness-datasource-inlinedocumentenrichmentconfiguration-condition
	//
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html#cfn-qbusiness-datasource-inlinedocumentenrichmentconfiguration-documentcontentoperator
	//
	DocumentContentOperator *string `field:"optional" json:"documentContentOperator" yaml:"documentContentOperator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html#cfn-qbusiness-datasource-inlinedocumentenrichmentconfiguration-target
	//
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

