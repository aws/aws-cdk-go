package awsram


// Properties for defining a `CfnPermissionAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissionAssociationProps := &CfnPermissionAssociationProps{
//   	PermissionArn: jsii.String("permissionArn"),
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//
//   	// the properties below are optional
//   	Replace: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-permissionassociation.html
//
type CfnPermissionAssociationProps struct {
	// Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the AWS RAM permission to associate with the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-permissionassociation.html#cfn-ram-permissionassociation-permissionarn
	//
	PermissionArn *string `field:"required" json:"permissionArn" yaml:"permissionArn"`
	// Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-permissionassociation.html#cfn-ram-permissionassociation-resourcesharearn
	//
	ResourceShareArn *string `field:"required" json:"resourceShareArn" yaml:"resourceShareArn"`
	// Specifies whether to replace the existing permission on the resource share.
	//
	// Use `true` to replace the current permission. Use `false` to add the permission when no permission is currently associated. The default value is `false`. Updating an existing association also requires `true`, because AWS RAM applies the change by re-associating the permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-permissionassociation.html#cfn-ram-permissionassociation-replace
	//
	// Default: - false.
	//
	Replace interface{} `field:"optional" json:"replace" yaml:"replace"`
}

