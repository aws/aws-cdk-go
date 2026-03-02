package previewawsconnectmixins


// Properties for CfnContactFlowModuleAliasPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContactFlowModuleAliasMixinProps := &CfnContactFlowModuleAliasMixinProps{
//   	ContactFlowModuleId: jsii.String("contactFlowModuleId"),
//   	ContactFlowModuleVersion: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html
//
type CfnContactFlowModuleAliasMixinProps struct {
	// The identifier of the contact flow module (ARN) this alias is tied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html#cfn-connect-contactflowmodulealias-contactflowmoduleid
	//
	ContactFlowModuleId *string `field:"optional" json:"contactFlowModuleId" yaml:"contactFlowModuleId"`
	// The version number of the contact flow module this alias points to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html#cfn-connect-contactflowmodulealias-contactflowmoduleversion
	//
	ContactFlowModuleVersion *float64 `field:"optional" json:"contactFlowModuleVersion" yaml:"contactFlowModuleVersion"`
	// The description of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html#cfn-connect-contactflowmodulealias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html#cfn-connect-contactflowmodulealias-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

