package awss3


// Specifies the details of Amazon S3 Storage Lens Group configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageLensGroupLevelProperty := &StorageLensGroupLevelProperty{
//   	StorageLensGroupSelectionCriteria: &StorageLensGroupSelectionCriteriaProperty{
//   		Exclude: []*string{
//   			jsii.String("exclude"),
//   		},
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgrouplevel.html
//
type CfnStorageLens_StorageLensGroupLevelProperty struct {
	// Selection criteria for Storage Lens Group level metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgrouplevel.html#cfn-s3-storagelens-storagelensgrouplevel-storagelensgroupselectioncriteria
	//
	StorageLensGroupSelectionCriteria interface{} `field:"optional" json:"storageLensGroupSelectionCriteria" yaml:"storageLensGroupSelectionCriteria"`
}

