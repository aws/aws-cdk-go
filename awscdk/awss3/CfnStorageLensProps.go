package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   cfnStorageLensProps := &CfnStorageLensProps{
//   	StorageLensConfiguration: &StorageLensConfigurationProperty{
//   		AccountLevel: &AccountLevelProperty{
//   			BucketLevel: &BucketLevelProperty{
//   				ActivityMetrics: &ActivityMetricsProperty{
//   					IsEnabled: jsii.Boolean(false),
//   				},
//   				AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   					IsEnabled: jsii.Boolean(false),
//   				},
//   				AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   					IsEnabled: jsii.Boolean(false),
//   				},
//   				DetailedStatusCodesMetrics: &DetailedStatusCodesMetricsProperty{
//   					IsEnabled: jsii.Boolean(false),
//   				},
//   				PrefixLevel: &PrefixLevelProperty{
//   					StorageMetrics: &PrefixLevelStorageMetricsProperty{
//   						IsEnabled: jsii.Boolean(false),
//   						SelectionCriteria: &SelectionCriteriaProperty{
//   							Delimiter: jsii.String("delimiter"),
//   							MaxDepth: jsii.Number(123),
//   							MinStorageBytesPercentage: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			ActivityMetrics: &ActivityMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			DetailedStatusCodesMetrics: &DetailedStatusCodesMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			StorageLensGroupLevel: &StorageLensGroupLevelProperty{
//   				StorageLensGroupSelectionCriteria: &StorageLensGroupSelectionCriteriaProperty{
//   					Exclude: []*string{
//   						jsii.String("exclude"),
//   					},
//   					Include: []*string{
//   						jsii.String("include"),
//   					},
//   				},
//   			},
//   		},
//   		Id: jsii.String("id"),
//   		IsEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AwsOrg: &AwsOrgProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		DataExport: &DataExportProperty{
//   			CloudWatchMetrics: &CloudWatchMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			S3BucketDestination: &S3BucketDestinationProperty{
//   				AccountId: jsii.String("accountId"),
//   				Arn: jsii.String("arn"),
//   				Format: jsii.String("format"),
//   				OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   				// the properties below are optional
//   				Encryption: &EncryptionProperty{
//   					Ssekms: &SSEKMSProperty{
//   						KeyId: jsii.String("keyId"),
//   					},
//   					Sses3: sses3,
//   				},
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   		Exclude: &BucketsAndRegionsProperty{
//   			Buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		Include: &BucketsAndRegionsProperty{
//   			Buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		StorageLensArn: jsii.String("storageLensArn"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html
//
type CfnStorageLensProps struct {
	// This resource contains the details Amazon S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html#cfn-s3-storagelens-storagelensconfiguration
	//
	StorageLensConfiguration interface{} `field:"required" json:"storageLensConfiguration" yaml:"storageLensConfiguration"`
	// A set of tags (key–value pairs) to associate with the Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html#cfn-s3-storagelens-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

