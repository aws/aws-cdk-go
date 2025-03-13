package awsquicksight


// Permission for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePermissionProperty := &ResourcePermissionProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	Principal: jsii.String("principal"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-resourcepermission.html
//
type CfnDataSet_ResourcePermissionProperty struct {
	// The IAM action to grant or revoke permisions on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-resourcepermission.html#cfn-quicksight-dataset-resourcepermission-actions
	//
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Name (ARN) of the principal. This can be one of the following:.
	//
	// - The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)
	// - The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)
	// - The ARN of an AWS account root: This is an IAM ARN rather than a Amazon QuickSight ARN. Use this option only to share resources (templates) across AWS accounts . (This is less common.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-resourcepermission.html#cfn-quicksight-dataset-resourcepermission-principal
	//
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

