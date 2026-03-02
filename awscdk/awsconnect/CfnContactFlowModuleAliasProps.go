package awsconnect


// Properties for defining a `CfnContactFlowModuleAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContactFlowModuleAliasProps := &CfnContactFlowModuleAliasProps{
//   	ContactFlowModuleId: jsii.String("contactFlowModuleId"),
//   	ContactFlowModuleVersion: jsii.Number(123),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html
//
type CfnContactFlowModuleAliasProps struct {
	// The identifier of the contact flow module (ARN) this alias is tied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html#cfn-connect-contactflowmodulealias-contactflowmoduleid
	//
	ContactFlowModuleId *string `field:"required" json:"contactFlowModuleId" yaml:"contactFlowModuleId"`
	// The version number of the contact flow module this alias points to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html#cfn-connect-contactflowmodulealias-contactflowmoduleversion
	//
	ContactFlowModuleVersion *float64 `field:"required" json:"contactFlowModuleVersion" yaml:"contactFlowModuleVersion"`
	// The name of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html#cfn-connect-contactflowmodulealias-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html#cfn-connect-contactflowmodulealias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

