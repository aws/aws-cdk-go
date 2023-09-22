package awsresiliencehub


// Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.
//
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
	// Defines how AWS Resilience Hub scans your resources.
	//
	// It can scan for the resources by using a pre-existing role in your AWS account, or by using the credentials of the current IAM user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-permissionmodel.html#cfn-resiliencehub-app-permissionmodel-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Defines a list of role Amazon Resource Names (ARNs) to be used in other accounts.
	//
	// These ARNs are used for querying purposes while importing resources and assessing your application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-permissionmodel.html#cfn-resiliencehub-app-permissionmodel-crossaccountrolearns
	//
	CrossAccountRoleArns *[]*string `field:"optional" json:"crossAccountRoleArns" yaml:"crossAccountRoleArns"`
	// Existing AWS IAM role name in the primary AWS account that will be assumed by AWS Resilience Hub Service Principle to obtain a read-only access to your application resources while running an assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-permissionmodel.html#cfn-resiliencehub-app-permissionmodel-invokerrolename
	//
	InvokerRoleName *string `field:"optional" json:"invokerRoleName" yaml:"invokerRoleName"`
}

