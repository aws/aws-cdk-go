package previewawss3mixins


// This is the property of the Amazon S3 Storage Lens configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var sses3 interface{}
//
//   storageLensConfigurationProperty := &StorageLensConfigurationProperty{
//   	AccountLevel: &AccountLevelProperty{
//   		ActivityMetrics: &ActivityMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		AdvancedPerformanceMetrics: &AdvancedPerformanceMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
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
//   			AdvancedPerformanceMetrics: &AdvancedPerformanceMetricsProperty{
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
//   			Encryption: &EncryptionProperty{
//   				Ssekms: &SSEKMSProperty{
//   					KeyId: jsii.String("keyId"),
//   				},
//   				Sses3: sses3,
//   			},
//   			Format: jsii.String("format"),
//   			OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		StorageLensTableDestination: &StorageLensTableDestinationProperty{
//   			Encryption: &EncryptionProperty{
//   				Ssekms: &SSEKMSProperty{
//   					KeyId: jsii.String("keyId"),
//   				},
//   				Sses3: sses3,
//   			},
//   			IsEnabled: jsii.Boolean(false),
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
//   	ExpandedPrefixesDataExport: &StorageLensExpandedPrefixesDataExportProperty{
//   		S3BucketDestination: &S3BucketDestinationProperty{
//   			AccountId: jsii.String("accountId"),
//   			Arn: jsii.String("arn"),
//   			Encryption: &EncryptionProperty{
//   				Ssekms: &SSEKMSProperty{
//   					KeyId: jsii.String("keyId"),
//   				},
//   				Sses3: sses3,
//   			},
//   			Format: jsii.String("format"),
//   			OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		StorageLensTableDestination: &StorageLensTableDestinationProperty{
//   			Encryption: &EncryptionProperty{
//   				Ssekms: &SSEKMSProperty{
//   					KeyId: jsii.String("keyId"),
//   				},
//   				Sses3: sses3,
//   			},
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	Include: &BucketsAndRegionsProperty{
//   		Buckets: []*string{
//   			jsii.String("buckets"),
//   		},
//   		Regions: []*string{
//   			jsii.String("regions"),
//   		},
//   	},
//   	IsEnabled: jsii.Boolean(false),
//   	PrefixDelimiter: jsii.String("prefixDelimiter"),
//   	StorageLensArn: jsii.String("storageLensArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html
//
type CfnStorageLensPropsMixin_StorageLensConfigurationProperty struct {
	// This property contains the details of the account-level metrics for Amazon S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-accountlevel
	//
	AccountLevel interface{} `field:"optional" json:"accountLevel" yaml:"accountLevel"`
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
	// Expanded Prefixes Data Export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-expandedprefixesdataexport
	//
	ExpandedPrefixesDataExport interface{} `field:"optional" json:"expandedPrefixesDataExport" yaml:"expandedPrefixesDataExport"`
	// This property contains the details of the ID of the S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// This property contains the details of the bucket and or Regions included for Amazon S3 Storage Lens configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-include
	//
	Include interface{} `field:"optional" json:"include" yaml:"include"`
	// This property contains the details of whether the Amazon S3 Storage Lens configuration is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-isenabled
	//
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// The delimiter to divide S3 key into hierarchy of prefixes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-prefixdelimiter
	//
	PrefixDelimiter *string `field:"optional" json:"prefixDelimiter" yaml:"prefixDelimiter"`
	// This property contains the details of the ARN of the S3 Storage Lens configuration.
	//
	// This property is read-only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensconfiguration.html#cfn-s3-storagelens-storagelensconfiguration-storagelensarn
	//
	StorageLensArn *string `field:"optional" json:"storageLensArn" yaml:"storageLensArn"`
}

