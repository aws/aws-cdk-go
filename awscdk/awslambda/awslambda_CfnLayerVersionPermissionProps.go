package awslambda


// Properties for defining a `CfnLayerVersionPermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLayerVersionPermissionProps := &cfnLayerVersionPermissionProps{
//   	action: jsii.String("action"),
//   	layerVersionArn: jsii.String("layerVersionArn"),
//   	principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	organizationId: jsii.String("organizationId"),
//   }
//
type CfnLayerVersionPermissionProps struct {
	// The API action that grants access to the layer.
	//
	// For example, `lambda:GetLayerVersion` .
	Action *string `field:"required" json:"action" yaml:"action"`
	// The name or Amazon Resource Name (ARN) of the layer.
	LayerVersionArn *string `field:"required" json:"layerVersionArn" yaml:"layerVersionArn"`
	// An account ID, or `*` to grant layer usage permission to all accounts in an organization, or all AWS accounts (if `organizationId` is not specified).
	//
	// For the last case, make sure that you really do want all AWS accounts to have usage permission to this layer.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// With the principal set to `*` , grant permission to all accounts in the specified organization.
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
}

