package mixinsawsbackup


// `RecoveryPointSelection` has five parameters (three required and two optional).
//
// The values you specify determine which recovery point is included in the restore test. You must indicate with `Algorithm` if you want the latest recovery point within your `SelectionWindowDays` or if you want a random recovery point, and you must indicate through `IncludeVaults` from which vaults the recovery points can be chosen.
//
// `Algorithm` ( *required* ) Valid values: " `LATEST_WITHIN_WINDOW` " or " `RANDOM_WITHIN_WINDOW` ".
//
// `Recovery point types` ( *required* ) Valid values: " `SNAPSHOT` " and/or " `CONTINUOUS` ". Include `SNAPSHOT` to restore only snapshot recovery points; include `CONTINUOUS` to restore continuous recovery points (point in time restore / PITR); use both to restore either a snapshot or a continuous recovery point. The recovery point will be determined by the value for `Algorithm` .
//
// `IncludeVaults` ( *required* ). You must include one or more backup vaults. Use the wildcard ["*"] or specific ARNs.
//
// `SelectionWindowDays` ( *optional* ) Value must be an integer (in days) from 1 to 365. If not included, the value defaults to `30` .
//
// `ExcludeVaults` ( *optional* ). You can choose to input one or more specific backup vault ARNs to exclude those vaults' contents from restore eligibility. Or, you can include a list of selectors. If this parameter and its value are not included, it defaults to empty list.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   restoreTestingRecoveryPointSelectionProperty := &RestoreTestingRecoveryPointSelectionProperty{
//   	Algorithm: jsii.String("algorithm"),
//   	ExcludeVaults: []*string{
//   		jsii.String("excludeVaults"),
//   	},
//   	IncludeVaults: []*string{
//   		jsii.String("includeVaults"),
//   	},
//   	RecoveryPointTypes: []*string{
//   		jsii.String("recoveryPointTypes"),
//   	},
//   	SelectionWindowDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html
//
type CfnRestoreTestingPlanPropsMixin_RestoreTestingRecoveryPointSelectionProperty struct {
	// Acceptable values include "LATEST_WITHIN_WINDOW" or "RANDOM_WITHIN_WINDOW".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-algorithm
	//
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Accepted values include specific ARNs or list of selectors.
	//
	// Defaults to empty list if not listed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-excludevaults
	//
	ExcludeVaults *[]*string `field:"optional" json:"excludeVaults" yaml:"excludeVaults"`
	// Accepted values include wildcard ["*"] or by specific ARNs or ARN wilcard replacement ["arn:aws:backup:us-west-2:123456789012:backup-vault:asdf", ...] ["arn:aws:backup:*:*:backup-vault:asdf-*", ...].
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-includevaults
	//
	IncludeVaults *[]*string `field:"optional" json:"includeVaults" yaml:"includeVaults"`
	// These are the types of recovery points.
	//
	// Include `SNAPSHOT` to restore only snapshot recovery points; include `CONTINUOUS` to restore continuous recovery points (point in time restore / PITR); use both to restore either a snapshot or a continuous recovery point. The recovery point will be determined by the value for `Algorithm` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-recoverypointtypes
	//
	RecoveryPointTypes *[]*string `field:"optional" json:"recoveryPointTypes" yaml:"recoveryPointTypes"`
	// Accepted values are integers from 1 to 365.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.html#cfn-backup-restoretestingplan-restoretestingrecoverypointselection-selectionwindowdays
	//
	SelectionWindowDays *float64 `field:"optional" json:"selectionWindowDays" yaml:"selectionWindowDays"`
}

