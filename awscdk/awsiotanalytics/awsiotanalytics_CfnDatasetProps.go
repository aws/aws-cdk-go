package awsiotanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetProps := &cfnDatasetProps{
//   	actions: []interface{}{
//   		&actionProperty{
//   			actionName: jsii.String("actionName"),
//
//   			// the properties below are optional
//   			containerAction: &containerActionProperty{
//   				executionRoleArn: jsii.String("executionRoleArn"),
//   				image: jsii.String("image"),
//   				resourceConfiguration: &resourceConfigurationProperty{
//   					computeType: jsii.String("computeType"),
//   					volumeSizeInGb: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				variables: []interface{}{
//   					&variableProperty{
//   						variableName: jsii.String("variableName"),
//
//   						// the properties below are optional
//   						datasetContentVersionValue: &datasetContentVersionValueProperty{
//   							datasetName: jsii.String("datasetName"),
//   						},
//   						doubleValue: jsii.Number(123),
//   						outputFileUriValue: &outputFileUriValueProperty{
//   							fileName: jsii.String("fileName"),
//   						},
//   						stringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			queryAction: &queryActionProperty{
//   				sqlQuery: jsii.String("sqlQuery"),
//
//   				// the properties below are optional
//   				filters: []interface{}{
//   					&filterProperty{
//   						deltaTime: &deltaTimeProperty{
//   							offsetSeconds: jsii.Number(123),
//   							timeExpression: jsii.String("timeExpression"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	contentDeliveryRules: []interface{}{
//   		&datasetContentDeliveryRuleProperty{
//   			destination: &datasetContentDeliveryRuleDestinationProperty{
//   				iotEventsDestinationConfiguration: &iotEventsDestinationConfigurationProperty{
//   					inputName: jsii.String("inputName"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				s3DestinationConfiguration: &s3DestinationConfigurationProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					glueConfiguration: &glueConfigurationProperty{
//   						databaseName: jsii.String("databaseName"),
//   						tableName: jsii.String("tableName"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			entryName: jsii.String("entryName"),
//   		},
//   	},
//   	datasetName: jsii.String("datasetName"),
//   	lateDataRules: []interface{}{
//   		&lateDataRuleProperty{
//   			ruleConfiguration: &lateDataRuleConfigurationProperty{
//   				deltaTimeSessionWindowConfiguration: &deltaTimeSessionWindowConfigurationProperty{
//   					timeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ruleName: jsii.String("ruleName"),
//   		},
//   	},
//   	retentionPeriod: &retentionPeriodProperty{
//   		numberOfDays: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	triggers: []interface{}{
//   		&triggerProperty{
//   			schedule: &scheduleProperty{
//   				scheduleExpression: jsii.String("scheduleExpression"),
//   			},
//   			triggeringDataset: &triggeringDatasetProperty{
//   				datasetName: jsii.String("datasetName"),
//   			},
//   		},
//   	},
//   	versioningConfiguration: &versioningConfigurationProperty{
//   		maxVersions: jsii.Number(123),
//   		unlimited: jsii.Boolean(false),
//   	},
//   }
//
type CfnDatasetProps struct {
	// The `DatasetAction` objects that automatically create the dataset contents.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// When dataset contents are created they are delivered to destinations specified here.
	ContentDeliveryRules interface{} `field:"optional" json:"contentDeliveryRules" yaml:"contentDeliveryRules"`
	// The name of the dataset.
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// A list of data rules that send notifications to CloudWatch, when data arrives late.
	//
	// To specify `lateDataRules` , the dataset must use a [DeltaTimer](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html) filter.
	LateDataRules interface{} `field:"optional" json:"lateDataRules" yaml:"lateDataRules"`
	// Optional.
	//
	// How long, in days, message data is kept for the dataset.
	RetentionPeriod interface{} `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Metadata which can be used to manage the data set.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The `DatasetTrigger` objects that specify when the dataset is automatically updated.
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// Optional.
	//
	// How many versions of dataset contents are kept. If not specified or set to null, only the latest version plus the latest succeeded version (if they are different) are kept for the time period specified by the `retentionPeriod` parameter. For more information, see [Keeping Multiple Versions of AWS IoT Analytics datasets](https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions) in the *AWS IoT Analytics User Guide* .
	VersioningConfiguration interface{} `field:"optional" json:"versioningConfiguration" yaml:"versioningConfiguration"`
}

