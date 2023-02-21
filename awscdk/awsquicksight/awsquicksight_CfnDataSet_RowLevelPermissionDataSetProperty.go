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
//   rowLevelPermissionDataSetProperty := &rowLevelPermissionDataSetProperty{
//   	arn: jsii.String("arn"),
//   	permissionPolicy: jsii.String("permissionPolicy"),
//
//   	// the properties below are optional
//   	formatVersion: jsii.String("formatVersion"),
//   	namespace: jsii.String("namespace"),
//   }
//
type CfnDataSet_RowLevelPermissionDataSetProperty struct {
	// The Amazon Resource Name (ARN) of the dataset that contains permissions for RLS.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The type of permissions to use when interpreting the permissions for RLS.
	//
	// `DENY_ACCESS` is included for backward compatibility only.
	PermissionPolicy *string `field:"required" json:"permissionPolicy" yaml:"permissionPolicy"`
	// The user or group rules associated with the dataset that contains permissions for RLS.
	//
	// By default, `FormatVersion` is `VERSION_1` . When `FormatVersion` is `VERSION_1` , `UserName` and `GroupName` are required. When `FormatVersion` is `VERSION_2` , `UserARN` and `GroupARN` are required, and `Namespace` must not exist.
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// The namespace associated with the dataset that contains permissions for RLS.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

