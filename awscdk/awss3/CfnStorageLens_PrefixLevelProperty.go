package awss3


// This resource contains the details of the prefix-level of the Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixLevelProperty := &PrefixLevelProperty{
//   	StorageMetrics: &PrefixLevelStorageMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   		SelectionCriteria: &SelectionCriteriaProperty{
//   			Delimiter: jsii.String("delimiter"),
//   			MaxDepth: jsii.Number(123),
//   			MinStorageBytesPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-prefixlevel.html
//
type CfnStorageLens_PrefixLevelProperty struct {
	// A property for the prefix-level storage metrics for Amazon S3 Storage Lens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-prefixlevel.html#cfn-s3-storagelens-prefixlevel-storagemetrics
	//
	StorageMetrics interface{} `field:"required" json:"storageMetrics" yaml:"storageMetrics"`
}

