package awsbackup


// Required: Algorithm;
//
// Required: Recovery point types; IncludeVaults(one or more). Optional: SelectionWindowDays ('30' if not specified);ExcludeVaults (list of selectors), defaults to empty list if not listed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restoreTestingRecoveryPointSelectionProperty := &RestoreTestingRecoveryPointSelectionProperty{
//   	Algorithm: jsii.String("algorithm"),
//   	IncludeVaults: []*string{
//   		jsii.String("includeVaults"),
//   	},
//   	RecoveryPointTypes: []*string{
//   		jsii.String("recoveryPointTypes"),
//   	},
//
//   	// the properties below are optional
//   	ExcludeVaults: []*string{
//   		jsii.String("excludeVaults"),
//   	},
//   	SelectionWindowDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html
//
type CfnRestoreTestingPlan_RestoreTestingRecoveryPointSelectionProperty struct {
	// Acceptable values include "LATEST_WITHIN_WINDOW" or "RANDOM_WITHIN_WINDOW".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-algorithm
	//
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Accepted values include wildcard ["*"] or by specific ARNs or ARN wilcard replacement ["arn:aws:backup:us-west-2:123456789012:backup-vault:asdf", ...] ["arn:aws:backup:*:*:backup-vault:asdf-*", ...].
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-includevaults
	//
	IncludeVaults *[]*string `field:"required" json:"includeVaults" yaml:"includeVaults"`
	// These are the types of recovery points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-recoverypointtypes
	//
	RecoveryPointTypes *[]*string `field:"required" json:"recoveryPointTypes" yaml:"recoveryPointTypes"`
	// Accepted values include specific ARNs or list of selectors.
	//
	// Defaults to empty list if not listed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-excludevaults
	//
	ExcludeVaults *[]*string `field:"optional" json:"excludeVaults" yaml:"excludeVaults"`
	// Accepted values are integers from 1 to 365.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-selectionwindowdays
	//
	SelectionWindowDays *float64 `field:"optional" json:"selectionWindowDays" yaml:"selectionWindowDays"`
}

