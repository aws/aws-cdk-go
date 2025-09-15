package awss3


// This resource indicates which Storage Lens group ARNs to include or exclude in the Storage Lens group aggregation.
//
// You can only attach Storage Lens groups to your dashboard if they're included in your Storage Lens group aggregation. If this value is left null, then all Storage Lens groups are selected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageLensGroupSelectionCriteriaProperty := &StorageLensGroupSelectionCriteriaProperty{
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	Include: []*string{
//   		jsii.String("include"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgroupselectioncriteria.html
//
type CfnStorageLens_StorageLensGroupSelectionCriteriaProperty struct {
	// This property indicates which Storage Lens group ARNs to exclude from the Storage Lens group aggregation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgroupselectioncriteria.html#cfn-s3-storagelens-storagelensgroupselectioncriteria-exclude
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// This property indicates which Storage Lens group ARNs to include in the Storage Lens group aggregation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgroupselectioncriteria.html#cfn-s3-storagelens-storagelensgroupselectioncriteria-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

