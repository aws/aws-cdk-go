package awsiotanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetProps := &CfnDatasetProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			ActionName: jsii.String("actionName"),
//
//   			// the properties below are optional
//   			ContainerAction: &ContainerActionProperty{
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				Image: jsii.String("image"),
//   				ResourceConfiguration: &ResourceConfigurationProperty{
//   					ComputeType: jsii.String("computeType"),
//   					VolumeSizeInGb: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				Variables: []interface{}{
//   					&VariableProperty{
//   						VariableName: jsii.String("variableName"),
//
//   						// the properties below are optional
//   						DatasetContentVersionValue: &DatasetContentVersionValueProperty{
//   							DatasetName: jsii.String("datasetName"),
//   						},
//   						DoubleValue: jsii.Number(123),
//   						OutputFileUriValue: &OutputFileUriValueProperty{
//   							FileName: jsii.String("fileName"),
//   						},
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			QueryAction: &QueryActionProperty{
//   				SqlQuery: jsii.String("sqlQuery"),
//
//   				// the properties below are optional
//   				Filters: []interface{}{
//   					&FilterProperty{
//   						DeltaTime: &DeltaTimeProperty{
//   							OffsetSeconds: jsii.Number(123),
//   							TimeExpression: jsii.String("timeExpression"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ContentDeliveryRules: []interface{}{
//   		&DatasetContentDeliveryRuleProperty{
//   			Destination: &DatasetContentDeliveryRuleDestinationProperty{
//   				IotEventsDestinationConfiguration: &IotEventsDestinationConfigurationProperty{
//   					InputName: jsii.String("inputName"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   				S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   					RoleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					GlueConfiguration: &GlueConfigurationProperty{
//   						DatabaseName: jsii.String("databaseName"),
//   						TableName: jsii.String("tableName"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			EntryName: jsii.String("entryName"),
//   		},
//   	},
//   	DatasetName: jsii.String("datasetName"),
//   	LateDataRules: []interface{}{
//   		&LateDataRuleProperty{
//   			RuleConfiguration: &LateDataRuleConfigurationProperty{
//   				DeltaTimeSessionWindowConfiguration: &DeltaTimeSessionWindowConfigurationProperty{
//   					TimeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			RuleName: jsii.String("ruleName"),
//   		},
//   	},
//   	RetentionPeriod: &RetentionPeriodProperty{
//   		NumberOfDays: jsii.Number(123),
//   		Unlimited: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Triggers: []interface{}{
//   		&TriggerProperty{
//   			Schedule: &ScheduleProperty{
//   				ScheduleExpression: jsii.String("scheduleExpression"),
//   			},
//   			TriggeringDataset: &TriggeringDatasetProperty{
//   				DatasetName: jsii.String("datasetName"),
//   			},
//   		},
//   	},
//   	VersioningConfiguration: &VersioningConfigurationProperty{
//   		MaxVersions: jsii.Number(123),
//   		Unlimited: jsii.Boolean(false),
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

