package awss3


// This resource contains the details of the prefix-level of the Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixLevelProperty := &prefixLevelProperty{
//   	storageMetrics: &prefixLevelStorageMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   		selectionCriteria: &selectionCriteriaProperty{
//   			delimiter: jsii.String("delimiter"),
//   			maxDepth: jsii.Number(123),
//   			minStorageBytesPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnStorageLens_PrefixLevelProperty struct {
	// A property for the prefix-level storage metrics for Amazon S3 Storage Lens.
	StorageMetrics interface{} `field:"required" json:"storageMetrics" yaml:"storageMetrics"`
}

