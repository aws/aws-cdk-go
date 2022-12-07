package awss3


// A property for the bucket-level storage metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketLevelProperty := &bucketLevelProperty{
//   	activityMetrics: &activityMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   	},
//   	advancedCostOptimizationMetrics: &advancedCostOptimizationMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   	},
//   	advancedDataProtectionMetrics: &advancedDataProtectionMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   	},
//   	detailedStatusCodesMetrics: &detailedStatusCodesMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   	},
//   	prefixLevel: &prefixLevelProperty{
//   		storageMetrics: &prefixLevelStorageMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   			selectionCriteria: &selectionCriteriaProperty{
//   				delimiter: jsii.String("delimiter"),
//   				maxDepth: jsii.Number(123),
//   				minStorageBytesPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnStorageLens_BucketLevelProperty struct {
	// A property for bucket-level activity metrics for S3 Storage Lens.
	ActivityMetrics interface{} `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// A property for bucket-level advanced cost optimization metrics for S3 Storage Lens.
	AdvancedCostOptimizationMetrics interface{} `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// A property for bucket-level advanced data protection metrics for S3 Storage Lens.
	AdvancedDataProtectionMetrics interface{} `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// A property for bucket-level detailed status code metrics for S3 Storage Lens.
	DetailedStatusCodesMetrics interface{} `field:"optional" json:"detailedStatusCodesMetrics" yaml:"detailedStatusCodesMetrics"`
	// A property for bucket-level prefix-level storage metrics for S3 Storage Lens.
	PrefixLevel interface{} `field:"optional" json:"prefixLevel" yaml:"prefixLevel"`
}

