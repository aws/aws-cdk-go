package awss3


// This resource determines the scope of Storage Lens group data that is displayed in the Storage Lens dashboard.
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
	// This property indicates which Storage Lens group ARNs to include or exclude in the Storage Lens group aggregation.
	//
	// If this value is left null, then all Storage Lens groups are selected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgrouplevel.html#cfn-s3-storagelens-storagelensgrouplevel-storagelensgroupselectioncriteria
	//
	StorageLensGroupSelectionCriteria interface{} `field:"optional" json:"storageLensGroupSelectionCriteria" yaml:"storageLensGroupSelectionCriteria"`
}

