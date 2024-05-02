package awsqbusiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   cfnDataSourceProps := &CfnDataSourceProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Configuration: configuration,
//   	DisplayName: jsii.String("displayName"),
//   	IndexId: jsii.String("indexId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DocumentEnrichmentConfiguration: &DocumentEnrichmentConfigurationProperty{
//   		InlineConfigurations: []interface{}{
//   			&InlineDocumentEnrichmentConfigurationProperty{
//   				Condition: &DocumentAttributeConditionProperty{
//   					Key: jsii.String("key"),
//   					Operator: jsii.String("operator"),
//
//   					// the properties below are optional
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
//   					Key: jsii.String("key"),
//
//   					// the properties below are optional
//   					AttributeValueOperator: jsii.String("attributeValueOperator"),
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
//   			LambdaArn: jsii.String("lambdaArn"),
//   			RoleArn: jsii.String("roleArn"),
//   			S3BucketName: jsii.String("s3BucketName"),
//   		},
//   		PreExtractionHookConfiguration: &HookConfigurationProperty{
//   			InvocationCondition: &DocumentAttributeConditionProperty{
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
//   			LambdaArn: jsii.String("lambdaArn"),
//   			RoleArn: jsii.String("roleArn"),
//   			S3BucketName: jsii.String("s3BucketName"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	SyncSchedule: jsii.String("syncSchedule"),
//   	Tags: []cfnTag{
//   		&cfnTag{
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
type CfnDataSourceProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-indexid
	//
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-documentenrichmentconfiguration
	//
	DocumentEnrichmentConfiguration interface{} `field:"optional" json:"documentEnrichmentConfiguration" yaml:"documentEnrichmentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-syncschedule
	//
	SyncSchedule *string `field:"optional" json:"syncSchedule" yaml:"syncSchedule"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-datasource.html#cfn-qbusiness-datasource-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

