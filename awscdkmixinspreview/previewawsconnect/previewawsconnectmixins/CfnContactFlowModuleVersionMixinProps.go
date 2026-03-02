package previewawsconnectmixins


// Properties for CfnContactFlowModuleVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContactFlowModuleVersionMixinProps := &CfnContactFlowModuleVersionMixinProps{
//   	ContactFlowModuleId: jsii.String("contactFlowModuleId"),
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmoduleversion.html
//
type CfnContactFlowModuleVersionMixinProps struct {
	// The identifier of the contact flow module (ARN) this version is tied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmoduleversion.html#cfn-connect-contactflowmoduleversion-contactflowmoduleid
	//
	ContactFlowModuleId *string `field:"optional" json:"contactFlowModuleId" yaml:"contactFlowModuleId"`
	// The description of the version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmoduleversion.html#cfn-connect-contactflowmoduleversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

