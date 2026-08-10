package awsquicksight


// A permission granted to a principal on a QuickSight resource.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-space-resourcepermission.html
//
type CfnSpace_ResourcePermissionProperty struct {
	// The list of actions granted to the principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-space-resourcepermission.html#cfn-quicksight-space-resourcepermission-actions
	//
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The ARN of the principal (user or group) receiving the permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-space-resourcepermission.html#cfn-quicksight-space-resourcepermission-principal
	//
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

