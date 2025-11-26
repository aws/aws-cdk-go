package previewawsquicksightmixins


// Permission for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourcePermissionProperty := &ResourcePermissionProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	Principal: jsii.String("principal"),
//   	Resource: jsii.String("resource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-resourcepermission.html
//
type CfnDashboardPropsMixin_ResourcePermissionProperty struct {
	// The IAM action to grant or revoke permissions on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-resourcepermission.html#cfn-quicksight-dashboard-resourcepermission-actions
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// The Amazon Resource Name (ARN) of the principal. This can be one of the following:.
	//
	// - The ARN of an Amazon Quick Suite user or group associated with a data source or dataset. (This is common.)
	// - The ARN of an Amazon Quick Suite user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)
	// - The ARN of an AWS account root: This is an IAM ARN rather than a Quick Suite ARN. Use this option only to share resources (templates) across AWS accounts . (This is less common.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-resourcepermission.html#cfn-quicksight-dashboard-resourcepermission-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-resourcepermission.html#cfn-quicksight-dashboard-resourcepermission-resource
	//
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

