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
//   storageLensConfigurationProperty := &storageLensConfigurationProperty{
//   	accountLevel: &accountLevelProperty{
//   		bucketLevel: &bucketLevelProperty{
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
//   			prefixLevel: &prefixLevelProperty{
//   				storageMetrics: &prefixLevelStorageMetricsProperty{
//   					isEnabled: jsii.Boolean(false),
//   					selectionCriteria: &selectionCriteriaProperty{
//   						delimiter: jsii.String("delimiter"),
//   						maxDepth: jsii.Number(123),
//   						minStorageBytesPercentage: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		activityMetrics: &activityMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   		},
//   		advancedCostOptimizationMetrics: &advancedCostOptimizationMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   		},
//   		advancedDataProtectionMetrics: &advancedDataProtectionMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   		},
//   		detailedStatusCodesMetrics: &detailedStatusCodesMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	id: jsii.String("id"),
//   	isEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	awsOrg: &awsOrgProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	dataExport: &dataExportProperty{
//   		cloudWatchMetrics: &cloudWatchMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   		},
//   		s3BucketDestination: &s3BucketDestinationProperty{
//   			accountId: jsii.String("accountId"),
//   			arn: jsii.String("arn"),
//   			format: jsii.String("format"),
//   			outputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   			// the properties below are optional
//   			encryption: &encryptionProperty{
//   				ssekms: &sSEKMSProperty{
//   					keyId: jsii.String("keyId"),
//   				},
//   				sses3: sses3,
//   			},
//   			prefix: jsii.String("prefix"),
//   		},
//   	},
//   	exclude: &bucketsAndRegionsProperty{
//   		buckets: []*string{
//   			jsii.String("buckets"),
//   		},
//   		regions: []*string{
//   			jsii.String("regions"),
//   		},
//   	},
//   	include: &bucketsAndRegionsProperty{
//   		buckets: []*string{
//   			jsii.String("buckets"),
//   		},
//   		regions: []*string{
//   			jsii.String("regions"),
//   		},
//   	},
//   	storageLensArn: jsii.String("storageLensArn"),
//   }
//
type CfnStorageLens_StorageLensConfigurationProperty struct {
	// This property contains the details of the account-level metrics for Amazon S3 Storage Lens configuration.
	AccountLevel interface{} `field:"required" json:"accountLevel" yaml:"accountLevel"`
	// This property contains the details of the ID of the S3 Storage Lens configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// This property contains the details of whether the Amazon S3 Storage Lens configuration is enabled.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// This property contains the details of the AWS Organization for the S3 Storage Lens configuration.
	AwsOrg interface{} `field:"optional" json:"awsOrg" yaml:"awsOrg"`
	// This property contains the details of this S3 Storage Lens configuration's metrics export.
	DataExport interface{} `field:"optional" json:"dataExport" yaml:"dataExport"`
	// This property contains the details of the bucket and or Regions excluded for Amazon S3 Storage Lens configuration.
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// This property contains the details of the bucket and or Regions included for Amazon S3 Storage Lens configuration.
	Include interface{} `field:"optional" json:"include" yaml:"include"`
	// This property contains the details of the ARN of the S3 Storage Lens configuration.
	//
	// This property is read-only.
	StorageLensArn *string `field:"optional" json:"storageLensArn" yaml:"storageLensArn"`
}

