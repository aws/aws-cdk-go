package awsquicksight


// Information about a dataset that contains permissions for row-level security (RLS).
//
// The permissions dataset maps fields to users or groups. For more information, see [Using Row-Level Security (RLS) to Restrict Access to a Dataset](https://docs.aws.amazon.com/quicksight/latest/user/restrict-access-to-a-data-set-using-row-level-security.html) in the *Amazon QuickSight User Guide* .
//
// The option to deny permissions by setting `PermissionPolicy` to `DENY_ACCESS` is not supported for new RLS datasets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rowLevelPermissionDataSetProperty := &RowLevelPermissionDataSetProperty{
//   	Arn: jsii.String("arn"),
//   	PermissionPolicy: jsii.String("permissionPolicy"),
//
//   	// the properties below are optional
//   	FormatVersion: jsii.String("formatVersion"),
//   	Namespace: jsii.String("namespace"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiondataset.html
//
type CfnDataSet_RowLevelPermissionDataSetProperty struct {
	// The Amazon Resource Name (ARN) of the dataset that contains permissions for RLS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiondataset.html#cfn-quicksight-dataset-rowlevelpermissiondataset-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The type of permissions to use when interpreting the permissions for RLS.
	//
	// `DENY_ACCESS` is included for backward compatibility only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiondataset.html#cfn-quicksight-dataset-rowlevelpermissiondataset-permissionpolicy
	//
	PermissionPolicy *string `field:"required" json:"permissionPolicy" yaml:"permissionPolicy"`
	// The user or group rules associated with the dataset that contains permissions for RLS.
	//
	// By default, `FormatVersion` is `VERSION_1` . When `FormatVersion` is `VERSION_1` , `UserName` and `GroupName` are required. When `FormatVersion` is `VERSION_2` , `UserARN` and `GroupARN` are required, and `Namespace` must not exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiondataset.html#cfn-quicksight-dataset-rowlevelpermissiondataset-formatversion
	//
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// The namespace associated with the dataset that contains permissions for RLS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiondataset.html#cfn-quicksight-dataset-rowlevelpermissiondataset-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The status of the row-level security permission dataset.
	//
	// If enabled, the status is `ENABLED` . If disabled, the status is `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-rowlevelpermissiondataset.html#cfn-quicksight-dataset-rowlevelpermissiondataset-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

