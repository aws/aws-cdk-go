package awss3


// This resource contains the details of the prefix-level storage metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixLevelStorageMetricsProperty := &PrefixLevelStorageMetricsProperty{
//   	IsEnabled: jsii.Boolean(false),
//   	SelectionCriteria: &SelectionCriteriaProperty{
//   		Delimiter: jsii.String("delimiter"),
//   		MaxDepth: jsii.Number(123),
//   		MinStorageBytesPercentage: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-prefixlevelstoragemetrics.html
//
type CfnStorageLens_PrefixLevelStorageMetricsProperty struct {
	// This property identifies whether the details of the prefix-level storage metrics for S3 Storage Lens are enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-prefixlevelstoragemetrics.html#cfn-s3-storagelens-prefixlevelstoragemetrics-isenabled
	//
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// This property identifies whether the details of the prefix-level storage metrics for S3 Storage Lens are enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-prefixlevelstoragemetrics.html#cfn-s3-storagelens-prefixlevelstoragemetrics-selectioncriteria
	//
	SelectionCriteria interface{} `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}

