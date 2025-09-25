package awslambda


// Properties for defining a `CfnLayerVersionPermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLayerVersionPermissionProps := &CfnLayerVersionPermissionProps{
//   	Action: jsii.String("action"),
//   	LayerVersionArn: jsii.String("layerVersionArn"),
//   	Principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	OrganizationId: jsii.String("organizationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html
//
type CfnLayerVersionPermissionProps struct {
	// The API action that grants access to the layer.
	//
	// For example, `lambda:GetLayerVersion` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The name or Amazon Resource Name (ARN) of the layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-layerversionarn
	//
	LayerVersionArn *string `field:"required" json:"layerVersionArn" yaml:"layerVersionArn"`
	// An account ID, or `*` to grant layer usage permission to all accounts in an organization, or all AWS accounts (if `organizationId` is not specified).
	//
	// For the last case, make sure that you really do want all AWS accounts to have usage permission to this layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-principal
	//
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// With the principal set to `*` , grant permission to all accounts in the specified organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-organizationid
	//
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
}

