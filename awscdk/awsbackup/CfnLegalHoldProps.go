package awsbackup


// Properties for defining a `CfnLegalHold`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLegalHoldProps := &CfnLegalHoldProps{
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
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html
//
type CfnLegalHoldProps struct {
	// The description of the legal hold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html#cfn-backup-legalhold-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The criteria to assign a set of resources, such as resource types or backup vaults.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html#cfn-backup-legalhold-recoverypointselection
	//
	RecoveryPointSelection interface{} `field:"required" json:"recoveryPointSelection" yaml:"recoveryPointSelection"`
	// The title of the legal hold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html#cfn-backup-legalhold-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// Optional tags to include.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-legalhold.html#cfn-backup-legalhold-tags
	//
	Tags *[]*CfnLegalHold_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

