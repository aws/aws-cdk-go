package previewawss3mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStorageLensPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var sses3 interface{}
//
//   cfnStorageLensMixinProps := &CfnStorageLensMixinProps{
//   	StorageLensConfiguration: &StorageLensConfigurationProperty{
//   		AccountLevel: &AccountLevelProperty{
//   			ActivityMetrics: &ActivityMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   			AdvancedPerformanceMetrics: &AdvancedPerformanceMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   			},
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
//   				AdvancedPerformanceMetrics: &AdvancedPerformanceMetricsProperty{
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
//   				Encryption: &EncryptionProperty{
//   					Ssekms: &SSEKMSProperty{
//   						KeyId: jsii.String("keyId"),
//   					},
//   					Sses3: sses3,
//   				},
//   				Format: jsii.String("format"),
//   				OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			StorageLensTableDestination: &StorageLensTableDestinationProperty{
//   				Encryption: &EncryptionProperty{
//   					Ssekms: &SSEKMSProperty{
//   						KeyId: jsii.String("keyId"),
//   					},
//   					Sses3: sses3,
//   				},
//   				IsEnabled: jsii.Boolean(false),
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
//   		ExpandedPrefixesDataExport: &StorageLensExpandedPrefixesDataExportProperty{
//   			S3BucketDestination: &S3BucketDestinationProperty{
//   				AccountId: jsii.String("accountId"),
//   				Arn: jsii.String("arn"),
//   				Encryption: &EncryptionProperty{
//   					Ssekms: &SSEKMSProperty{
//   						KeyId: jsii.String("keyId"),
//   					},
//   					Sses3: sses3,
//   				},
//   				Format: jsii.String("format"),
//   				OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			StorageLensTableDestination: &StorageLensTableDestinationProperty{
//   				Encryption: &EncryptionProperty{
//   					Ssekms: &SSEKMSProperty{
//   						KeyId: jsii.String("keyId"),
//   					},
//   					Sses3: sses3,
//   				},
//   				IsEnabled: jsii.Boolean(false),
//   			},
//   		},
//   		Id: jsii.String("id"),
//   		Include: &BucketsAndRegionsProperty{
//   			Buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		IsEnabled: jsii.Boolean(false),
//   		PrefixDelimiter: jsii.String("prefixDelimiter"),
//   		StorageLensArn: jsii.String("storageLensArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html
//
type CfnStorageLensMixinProps struct {
	// This resource contains the details Amazon S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html#cfn-s3-storagelens-storagelensconfiguration
	//
	StorageLensConfiguration interface{} `field:"optional" json:"storageLensConfiguration" yaml:"storageLensConfiguration"`
	// A set of tags (key–value pairs) to associate with the Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-storagelens.html#cfn-s3-storagelens-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

