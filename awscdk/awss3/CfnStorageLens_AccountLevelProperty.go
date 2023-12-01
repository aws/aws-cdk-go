package awss3


// This resource contains the details of the account-level metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountLevelProperty := &AccountLevelProperty{
//   	BucketLevel: &BucketLevelProperty{
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
//   		PrefixLevel: &PrefixLevelProperty{
//   			StorageMetrics: &PrefixLevelStorageMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   				SelectionCriteria: &SelectionCriteriaProperty{
//   					Delimiter: jsii.String("delimiter"),
//   					MaxDepth: jsii.Number(123),
//   					MinStorageBytesPercentage: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ActivityMetrics: &ActivityMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   	AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   	AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   	DetailedStatusCodesMetrics: &DetailedStatusCodesMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   	StorageLensGroupLevel: &StorageLensGroupLevelProperty{
//   		StorageLensGroupSelectionCriteria: &StorageLensGroupSelectionCriteriaProperty{
//   			Exclude: []*string{
//   				jsii.String("exclude"),
//   			},
//   			Include: []*string{
//   				jsii.String("include"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-accountlevel.html
//
type CfnStorageLens_AccountLevelProperty struct {
	// This property contains the details of the account-level bucket-level configurations for Amazon S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-accountlevel.html#cfn-s3-storagelens-accountlevel-bucketlevel
	//
	BucketLevel interface{} `field:"required" json:"bucketLevel" yaml:"bucketLevel"`
	// This property contains the details of account-level activity metrics for S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-accountlevel.html#cfn-s3-storagelens-accountlevel-activitymetrics
	//
	ActivityMetrics interface{} `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// This property contains the details of account-level advanced cost optimization metrics for S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-accountlevel.html#cfn-s3-storagelens-accountlevel-advancedcostoptimizationmetrics
	//
	AdvancedCostOptimizationMetrics interface{} `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// This property contains the details of account-level advanced data protection metrics for S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-accountlevel.html#cfn-s3-storagelens-accountlevel-advanceddataprotectionmetrics
	//
	AdvancedDataProtectionMetrics interface{} `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// This property contains the details of account-level detailed status code metrics for S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-accountlevel.html#cfn-s3-storagelens-accountlevel-detailedstatuscodesmetrics
	//
	DetailedStatusCodesMetrics interface{} `field:"optional" json:"detailedStatusCodesMetrics" yaml:"detailedStatusCodesMetrics"`
	// This property determines the scope of Storage Lens group data that is displayed in the Storage Lens dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-accountlevel.html#cfn-s3-storagelens-accountlevel-storagelensgrouplevel
	//
	StorageLensGroupLevel interface{} `field:"optional" json:"storageLensGroupLevel" yaml:"storageLensGroupLevel"`
}

