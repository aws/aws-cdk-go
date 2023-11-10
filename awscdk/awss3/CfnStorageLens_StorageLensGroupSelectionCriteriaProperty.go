package awss3


// Selection criteria for Storage Lens Group level metrics.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgroupselectioncriteria.html#cfn-s3-storagelens-storagelensgroupselectioncriteria-exclude
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-storagelensgroupselectioncriteria.html#cfn-s3-storagelens-storagelensgroupselectioncriteria-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

