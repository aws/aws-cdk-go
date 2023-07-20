package awss3


// This resource contains the details of the Amazon S3 Storage Lens selection criteria.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectionCriteriaProperty := &SelectionCriteriaProperty{
//   	Delimiter: jsii.String("delimiter"),
//   	MaxDepth: jsii.Number(123),
//   	MinStorageBytesPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-selectioncriteria.html
//
type CfnStorageLens_SelectionCriteriaProperty struct {
	// This property contains the details of the S3 Storage Lens delimiter being used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-selectioncriteria.html#cfn-s3-storagelens-selectioncriteria-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// This property contains the details of the max depth that S3 Storage Lens will collect metrics up to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-selectioncriteria.html#cfn-s3-storagelens-selectioncriteria-maxdepth
	//
	MaxDepth *float64 `field:"optional" json:"maxDepth" yaml:"maxDepth"`
	// This property contains the details of the minimum storage bytes percentage threshold that S3 Storage Lens will collect metrics up to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-selectioncriteria.html#cfn-s3-storagelens-selectioncriteria-minstoragebytespercentage
	//
	MinStorageBytesPercentage *float64 `field:"optional" json:"minStorageBytesPercentage" yaml:"minStorageBytesPercentage"`
}

