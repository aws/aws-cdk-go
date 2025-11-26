package previewawsconnectmixins


// Properties for CfnContactFlowVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContactFlowVersionMixinProps := &CfnContactFlowVersionMixinProps{
//   	ContactFlowId: jsii.String("contactFlowId"),
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowversion.html
//
type CfnContactFlowVersionMixinProps struct {
	// The identifier of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowversion.html#cfn-connect-contactflowversion-contactflowid
	//
	ContactFlowId *string `field:"optional" json:"contactFlowId" yaml:"contactFlowId"`
	// The description of the flow version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowversion.html#cfn-connect-contactflowversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

