package previewawsqbusinessmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDataSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configuration interface{}
//
//   cfnDataSourceMixinProps := &CfnDataSourceMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Configuration: configuration,
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	DocumentEnrichmentConfiguration: &DocumentEnrichmentConfigurationProperty{
//   		InlineConfigurations: []interface{}{
//   			&InlineDocumentEnrichmentConfigurationProperty{
//   				Condition: &DocumentAttributeConditionProperty{
//   					Key: jsii.String("key"),
//   					Operator: jsii.String("operator"),
//   					Value: &DocumentAttributeValueProperty{
//   						DateValue: jsii.String("dateValue"),
//   						LongValue: jsii.Number(123),
//   						StringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   				DocumentContentOperator: jsii.String("documentContentOperator"),
//   				Target: &DocumentAttributeTargetProperty{
//   					AttributeValueOperator: jsii.String("attributeValueOperator"),
//   					Key: jsii.String("key"),
//   					Value: &DocumentAttributeValueProperty{
//   						DateValue: jsii.String("dateValue"),
//   						LongValue: jsii.Number(123),
//   						StringListValue: []*string{
//   							jsii.String("stringListValue"),
//   						},
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   		},
//   		PostExtractionHookConfiguration: &HookConfigurationProperty{
//   			InvocationCondition: &DocumentAttributeConditionProperty{
//   				Key: jsii.String("key"),
//   				Operator: jsii.String("operator"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			LambdaArn: jsii.String("lambdaArn"),
//   			RoleArn: jsii.String("roleArn"),
//   			S3BucketName: jsii.String("s3BucketName"),
//   		},
//   		PreExtractionHookConfiguration: &HookConfigurationProperty{
//   			InvocationCondition: &DocumentAttributeConditionProperty{
//   				Key: jsii.String("key"),
//   				Operator: jsii.String("operator"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			LambdaArn: jsii.String("lambdaArn"),
//   			RoleArn: jsii.String("roleArn"),
//   			S3BucketName: jsii.String("s3BucketName"),
//   		},
//   	},
//   	IndexId: jsii.String("indexId"),
//   	MediaExtractionConfiguration: &MediaExtractionConfigurationProperty{
//   		AudioExtractionConfiguration: &AudioExtractionConfigurationProperty{
//   			AudioExtractionStatus: jsii.String("audioExtractionStatus"),
//   		},
//   		ImageExtractionConfiguration: &ImageExtractionConfigurationProperty{
//   			ImageExtractionStatus: jsii.String("imageExtractionStatus"),
//   		},
//   		VideoExtractionConfiguration: &VideoExtractionConfigurationProperty{
//   			VideoExtractionStatus: jsii.String("videoExtractionStatus"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	SyncSchedule: jsii.String("syncSchedule"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html
//
type CfnDataSourceMixinProps struct {
	// The identifier of the Amazon Q Business application the data source will be attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Use this property to specify a JSON or YAML schema with configuration properties specific to your data source connector to connect your data source repository to Amazon Q Business .
	//
	// You must use the JSON or YAML schema provided by Amazon Q .
	//
	// The following links have the configuration properties and schemas for AWS CloudFormation for the following connectors:
	//
	// - [Amazon Simple Storage Service](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/s3-cfn.html)
	// - [Amazon Q Web Crawler](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/web-crawler-cfn.html)
	//
	// Similarly, you can find configuration templates and properties for your specific data source using the following steps:
	//
	// - Navigate to the [Supported connectors](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/connectors-list.html) page in the Amazon Q Business User Guide, and select the data source connector of your choice.
	// - Then, from that specific data source connector's page, choose the topic containing *Using CloudFormation* to find the schemas for your data source connector, including configuration parameter descriptions and examples.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A description for the data source connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the Amazon Q Business data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Provides the configuration information for altering document metadata and content during the document ingestion process.
	//
	// For more information, see [Custom document enrichment](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/custom-document-enrichment.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-documentenrichmentconfiguration
	//
	DocumentEnrichmentConfiguration interface{} `field:"optional" json:"documentEnrichmentConfiguration" yaml:"documentEnrichmentConfiguration"`
	// The identifier of the index the data source is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-indexid
	//
	IndexId *string `field:"optional" json:"indexId" yaml:"indexId"`
	// The configuration for extracting information from media in documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-mediaextractionconfiguration
	//
	MediaExtractionConfiguration interface{} `field:"optional" json:"mediaExtractionConfiguration" yaml:"mediaExtractionConfiguration"`
	// The Amazon Resource Name (ARN) of an IAM role with permission to access the data source and required resources.
	//
	// This field is required for all connector types except custom connectors, where it is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Sets the frequency for Amazon Q Business to check the documents in your data source repository and update your index.
	//
	// If you don't set a schedule, Amazon Q Business won't periodically update the index.
	//
	// Specify a `cron-` format schedule string or an empty string to indicate that the index is updated on demand. You can't specify the `Schedule` parameter when the `Type` parameter is set to `CUSTOM` . If you do, you receive a `ValidationException` exception.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-syncschedule
	//
	SyncSchedule *string `field:"optional" json:"syncSchedule" yaml:"syncSchedule"`
	// A list of key-value pairs that identify or categorize the data source connector.
	//
	// You can also use tags to help control access to the data source connector. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + -
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Configuration information for an Amazon VPC (Virtual Private Cloud) to connect to your data source.
	//
	// For more information, see [Using Amazon VPC with Amazon Q Business connectors](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/connector-vpc.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

