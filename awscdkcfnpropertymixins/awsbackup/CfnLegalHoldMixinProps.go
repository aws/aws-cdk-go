package awsbackup


// Properties for CfnLegalHoldPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnLegalHoldMixinProps := &CfnLegalHoldMixinProps{
//   	Description: jsii.String("description"),
//   	RecoveryPointSelection: &RecoveryPointSelectionProperty{
//   		DateRange: &DateRangeProperty{
//   			FromDate: jsii.String("fromDate"),
//   			ToDate: jsii.String("toDate"),
//   		},
//   		ResourceIdentifiers: []*string{
//   			jsii.String("resourceIdentifiers"),
//   		},
//   		VaultNames: []*string{
//   			jsii.String("vaultNames"),
//   		},
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html
//
type CfnLegalHoldMixinProps struct {
	// The description of the legal hold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html#cfn-backup-legalhold-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The criteria to assign a set of resources, such as resource types or backup vaults.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html#cfn-backup-legalhold-recoverypointselection
	//
	RecoveryPointSelection interface{} `field:"optional" json:"recoveryPointSelection" yaml:"recoveryPointSelection"`
	// Optional tags to include.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html#cfn-backup-legalhold-tags
	//
	Tags *[]*CfnLegalHoldPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// The title of the legal hold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html#cfn-backup-legalhold-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

