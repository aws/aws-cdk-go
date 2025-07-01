package awss3


// A property for the bucket-level storage metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketLevelProperty := &BucketLevelProperty{
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
//   	PrefixLevel: &PrefixLevelProperty{
//   		StorageMetrics: &PrefixLevelStorageMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   			SelectionCriteria: &SelectionCriteriaProperty{
//   				Delimiter: jsii.String("delimiter"),
//   				MaxDepth: jsii.Number(123),
//   				MinStorageBytesPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketlevel.html
//
type CfnStorageLens_BucketLevelProperty struct {
	// A property for bucket-level activity metrics for S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketlevel.html#cfn-s3-storagelens-bucketlevel-activitymetrics
	//
	ActivityMetrics interface{} `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// A property for bucket-level advanced cost optimization metrics for S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketlevel.html#cfn-s3-storagelens-bucketlevel-advancedcostoptimizationmetrics
	//
	AdvancedCostOptimizationMetrics interface{} `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// A property for bucket-level advanced data protection metrics for S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketlevel.html#cfn-s3-storagelens-bucketlevel-advanceddataprotectionmetrics
	//
	AdvancedDataProtectionMetrics interface{} `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// A property for bucket-level detailed status code metrics for S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketlevel.html#cfn-s3-storagelens-bucketlevel-detailedstatuscodesmetrics
	//
	DetailedStatusCodesMetrics interface{} `field:"optional" json:"detailedStatusCodesMetrics" yaml:"detailedStatusCodesMetrics"`
	// A property for bucket-level prefix-level storage metrics for S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketlevel.html#cfn-s3-storagelens-bucketlevel-prefixlevel
	//
	PrefixLevel interface{} `field:"optional" json:"prefixLevel" yaml:"prefixLevel"`
}

