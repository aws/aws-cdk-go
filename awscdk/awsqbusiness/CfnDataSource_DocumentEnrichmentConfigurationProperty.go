package awsqbusiness


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentEnrichmentConfigurationProperty := &DocumentEnrichmentConfigurationProperty{
//   	InlineConfigurations: []interface{}{
//   		&InlineDocumentEnrichmentConfigurationProperty{
//   			Condition: &DocumentAttributeConditionProperty{
//   				Key: jsii.String("key"),
//   				Operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			DocumentContentOperator: jsii.String("documentContentOperator"),
//   			Target: &DocumentAttributeTargetProperty{
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				AttributeValueOperator: jsii.String("attributeValueOperator"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   		},
//   	},
//   	PostExtractionHookConfiguration: &HookConfigurationProperty{
//   		InvocationCondition: &DocumentAttributeConditionProperty{
//   			Key: jsii.String("key"),
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		LambdaArn: jsii.String("lambdaArn"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   	},
//   	PreExtractionHookConfiguration: &HookConfigurationProperty{
//   		InvocationCondition: &DocumentAttributeConditionProperty{
//   			Key: jsii.String("key"),
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		LambdaArn: jsii.String("lambdaArn"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentenrichmentconfiguration.html
//
type CfnDataSource_DocumentEnrichmentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentenrichmentconfiguration.html#cfn-qbusiness-datasource-documentenrichmentconfiguration-inlineconfigurations
	//
	InlineConfigurations interface{} `field:"optional" json:"inlineConfigurations" yaml:"inlineConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentenrichmentconfiguration.html#cfn-qbusiness-datasource-documentenrichmentconfiguration-postextractionhookconfiguration
	//
	PostExtractionHookConfiguration interface{} `field:"optional" json:"postExtractionHookConfiguration" yaml:"postExtractionHookConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentenrichmentconfiguration.html#cfn-qbusiness-datasource-documentenrichmentconfiguration-preextractionhookconfiguration
	//
	PreExtractionHookConfiguration interface{} `field:"optional" json:"preExtractionHookConfiguration" yaml:"preExtractionHookConfiguration"`
}

