package previewawsquicksightmixins


// <p>Permission for the resource.</p>.
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-folder-resourcepermission.html
//
type CfnFolderPropsMixin_ResourcePermissionProperty struct {
	// <p>The IAM action to grant or revoke permissions on.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-folder-resourcepermission.html#cfn-quicksight-folder-resourcepermission-actions
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// <p>The Amazon Resource Name (ARN) of the principal.
	//
	// This can be one of the
	//             following:</p>
	//          <ul>
	//             <li>
	//                <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>
	//             </li>
	//             <li>
	//                <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>
	//             </li>
	//             <li>
	//                <p>The ARN of an Amazon Web Services account root: This is an IAM ARN rather than a QuickSight
	//                     ARN. Use this option only to share resources (templates) across Amazon Web Services accounts.
	//                     (This is less common.) </p>
	//             </li>
	// </ul>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-folder-resourcepermission.html#cfn-quicksight-folder-resourcepermission-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

