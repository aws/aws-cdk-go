package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStorageLens`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sses3 interface{}
//
//   cfnStorageLensProps := &cfnStorageLensProps{
//   	storageLensConfiguration: &storageLensConfigurationProperty{
//   		accountLevel: &accountLevelProperty{
//   			bucketLevel: &bucketLevelProperty{
//   				activityMetrics: &activityMetricsProperty{
//   					isEnabled: jsii.Boolean(false),
//   				},
//   				advancedCostOptimizationMetrics: &advancedCostOptimizationMetricsProperty{
//   					isEnabled: jsii.Boolean(false),
//   				},
//   				advancedDataProtectionMetrics: &advancedDataProtectionMetricsProperty{
//   					isEnabled: jsii.Boolean(false),
//   				},
//   				detailedStatusCodesMetrics: &detailedStatusCodesMetricsProperty{
//   					isEnabled: jsii.Boolean(false),
//   				},
//   				prefixLevel: &prefixLevelProperty{
//   					storageMetrics: &prefixLevelStorageMetricsProperty{
//   						isEnabled: jsii.Boolean(false),
//   						selectionCriteria: &selectionCriteriaProperty{
//   							delimiter: jsii.String("delimiter"),
//   							maxDepth: jsii.Number(123),
//   							minStorageBytesPercentage: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			activityMetrics: &activityMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   			advancedCostOptimizationMetrics: &advancedCostOptimizationMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   			advancedDataProtectionMetrics: &advancedDataProtectionMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   			detailedStatusCodesMetrics: &detailedStatusCodesMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   		},
//   		id: jsii.String("id"),
//   		isEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		awsOrg: &awsOrgProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		dataExport: &dataExportProperty{
//   			cloudWatchMetrics: &cloudWatchMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   			s3BucketDestination: &s3BucketDestinationProperty{
//   				accountId: jsii.String("accountId"),
//   				arn: jsii.String("arn"),
//   				format: jsii.String("format"),
//   				outputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   				// the properties below are optional
//   				encryption: &encryptionProperty{
//   					ssekms: &sSEKMSProperty{
//   						keyId: jsii.String("keyId"),
//   					},
//   					sses3: sses3,
//   				},
//   				prefix: jsii.String("prefix"),
//   			},
//   		},
//   		exclude: &bucketsAndRegionsProperty{
//   			buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		include: &bucketsAndRegionsProperty{
//   			buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		storageLensArn: jsii.String("storageLensArn"),
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStorageLensProps struct {
	// This resource contains the details Amazon S3 Storage Lens configuration.
	StorageLensConfiguration interface{} `field:"required" json:"storageLensConfiguration" yaml:"storageLensConfiguration"`
	// A set of tags (key–value pairs) to associate with the Storage Lens configuration.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

