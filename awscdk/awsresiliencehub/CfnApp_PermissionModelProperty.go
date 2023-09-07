package awsresiliencehub


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionModelProperty := &PermissionModelProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	CrossAccountRoleArns: []*string{
//   		jsii.String("crossAccountRoleArns"),
//   	},
//   	InvokerRoleName: jsii.String("invokerRoleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-permissionmodel.html
//
type CfnApp_PermissionModelProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-permissionmodel.html#cfn-resiliencehub-app-permissionmodel-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-permissionmodel.html#cfn-resiliencehub-app-permissionmodel-crossaccountrolearns
	//
	CrossAccountRoleArns *[]*string `field:"optional" json:"crossAccountRoleArns" yaml:"crossAccountRoleArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-permissionmodel.html#cfn-resiliencehub-app-permissionmodel-invokerrolename
	//
	InvokerRoleName *string `field:"optional" json:"invokerRoleName" yaml:"invokerRoleName"`
}

