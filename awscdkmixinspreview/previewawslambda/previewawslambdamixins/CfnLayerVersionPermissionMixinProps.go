package previewawslambdamixins


// Properties for CfnLayerVersionPermissionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLayerVersionPermissionMixinProps := &CfnLayerVersionPermissionMixinProps{
//   	Action: jsii.String("action"),
//   	LayerVersionArn: jsii.String("layerVersionArn"),
//   	OrganizationId: jsii.String("organizationId"),
//   	Principal: jsii.String("principal"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html
//
type CfnLayerVersionPermissionMixinProps struct {
	// The API action that grants access to the layer.
	//
	// For example, `lambda:GetLayerVersion` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The name or Amazon Resource Name (ARN) of the layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-layerversionarn
	//
	LayerVersionArn *string `field:"optional" json:"layerVersionArn" yaml:"layerVersionArn"`
	// With the principal set to `*` , grant permission to all accounts in the specified organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-organizationid
	//
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// An account ID, or `*` to grant layer usage permission to all accounts in an organization, or all AWS accounts (if `organizationId` is not specified).
	//
	// For the last case, make sure that you really do want all AWS accounts to have usage permission to this layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

