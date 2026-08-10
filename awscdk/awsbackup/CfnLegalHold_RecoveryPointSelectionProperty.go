package awsbackup


// The criteria to assign a set of resources, such as resource types or backup vaults.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recoveryPointSelectionProperty := &RecoveryPointSelectionProperty{
//   	DateRange: &DateRangeProperty{
//   		FromDate: jsii.String("fromDate"),
//   		ToDate: jsii.String("toDate"),
//   	},
//   	ResourceIdentifiers: []*string{
//   		jsii.String("resourceIdentifiers"),
//   	},
//   	VaultNames: []*string{
//   		jsii.String("vaultNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-recoverypointselection.html
//
type CfnLegalHold_RecoveryPointSelectionProperty struct {
	// A date range for filtering recovery points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-recoverypointselection.html#cfn-backup-legalhold-recoverypointselection-daterange
	//
	DateRange interface{} `field:"optional" json:"dateRange" yaml:"dateRange"`
	// The resources included in the resource selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-recoverypointselection.html#cfn-backup-legalhold-recoverypointselection-resourceidentifiers
	//
	ResourceIdentifiers *[]*string `field:"optional" json:"resourceIdentifiers" yaml:"resourceIdentifiers"`
	// The names of the vaults in which the selected recovery points are contained.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-recoverypointselection.html#cfn-backup-legalhold-recoverypointselection-vaultnames
	//
	VaultNames *[]*string `field:"optional" json:"vaultNames" yaml:"vaultNames"`
}

