package awss3


// This resource contains the details of the Amazon S3 Storage Lens selection criteria.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectionCriteriaProperty := &selectionCriteriaProperty{
//   	delimiter: jsii.String("delimiter"),
//   	maxDepth: jsii.Number(123),
//   	minStorageBytesPercentage: jsii.Number(123),
//   }
//
type CfnStorageLens_SelectionCriteriaProperty struct {
	// This property contains the details of the S3 Storage Lens delimiter being used.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// This property contains the details of the max depth that S3 Storage Lens will collect metrics up to.
	MaxDepth *float64 `field:"optional" json:"maxDepth" yaml:"maxDepth"`
	// This property contains the details of the minimum storage bytes percentage threshold that S3 Storage Lens will collect metrics up to.
	MinStorageBytesPercentage *float64 `field:"optional" json:"minStorageBytesPercentage" yaml:"minStorageBytesPercentage"`
}

