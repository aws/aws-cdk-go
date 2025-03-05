package awss3


// This is the property of the Amazon S3 Storage Lens configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sses3 interface{}
//
//   storageLensConfigurationProperty := &StorageLensConfigurationProperty{
//   	AccountLevel: &AccountLevelProperty{
//   		BucketLevel: &BucketLevelProperty{
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
//   			PrefixLevel: &PrefixLevelProperty{
//   				StorageMetrics: &PrefixLevelStorageMetricsProperty{
//   					IsEnabled: jsii.Boolean(false),
//   					SelectionCriteria: &SelectionCriteriaProperty{
//   						Delimiter: jsii.String("delimiter"),
//   						MaxDepth: jsii.Number(123),
//   						MinStorageBytesPercentage: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		ActivityMetrics: &ActivityMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		DetailedStatusCodesMetrics: &DetailedStatusCodesMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		StorageLensGroupLevel: &StorageLensGroupLevelProperty{
//   			StorageLensGroupSelectionCriteria: &StorageLensGroupSelectionCriteriaProperty{
//   				Exclude: []*string{
//   					jsii.String("exclude"),
//   				},
//   				Include: []*string{
//   					jsii.String("include"),
//   				},
//   			},
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	IsEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AwsOrg: &AwsOrgProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	DataExport: &DataExportProperty{
//   		CloudWatchMetrics: &CloudWatchMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		S3BucketDestination: &S3BucketDestinationProperty{
//   			AccountId: jsii.String("accountId"),
//   			Arn: jsii.String("arn"),
//   			Format: jsii.String("format"),
//   			OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   			// the properties below are optional
//   			Encryption: &EncryptionProperty{
//   				Ssekms: &SSEKMSProperty{
//   					KeyId: jsii.String("keyId"),
//   				},
//   				Sses3: sses3,
//   			},
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	Exclude: &BucketsAndRegionsProperty{
//   		Buckets: []*string{
//   			jsii.String("buckets"),
//   		},
//   		Regions: []*string{
//   			jsii.String("regions"),
//   		},
//   	},
//   	Include: &BucketsAndRegionsProperty{
//   		Buckets: []*string{
//   			jsii.String("buckets"),
//   		},
//   		Regions: []*string{
//   			jsii.String("regions"),
//   		},
//   	},
//   	StorageLensArn: jsii.String("storageLensArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html
//
type CfnStorageLens_StorageLensConfigurationProperty struct {
	// This property contains the details of the account-level metrics for Amazon S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-accountlevel
	//
	AccountLevel interface{} `field:"required" json:"accountLevel" yaml:"accountLevel"`
	// This property contains the details of the ID of the S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// This property contains the details of whether the Amazon S3 Storage Lens configuration is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-isenabled
	//
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// This property contains the details of the AWS Organization for the S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-awsorg
	//
	AwsOrg interface{} `field:"optional" json:"awsOrg" yaml:"awsOrg"`
	// This property contains the details of this S3 Storage Lens configuration's metrics export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-dataexport
	//
	DataExport interface{} `field:"optional" json:"dataExport" yaml:"dataExport"`
	// This property contains the details of the bucket and or Regions excluded for Amazon S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-exclude
	//
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// This property contains the details of the bucket and or Regions included for Amazon S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-include
	//
	Include interface{} `field:"optional" json:"include" yaml:"include"`
	// This property contains the details of the ARN of the S3 Storage Lens configuration.
	//
	// This property is read-only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-storagelensarn
	//
	StorageLensArn *string `field:"optional" json:"storageLensArn" yaml:"storageLensArn"`
}

