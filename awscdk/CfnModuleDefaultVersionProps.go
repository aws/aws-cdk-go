package awscdk


// Properties for defining a `CfnModuleDefaultVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModuleDefaultVersionProps := &CfnModuleDefaultVersionProps{
//   	Arn: jsii.String("arn"),
//   	ModuleName: jsii.String("moduleName"),
//   	VersionId: jsii.String("versionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduledefaultversion.html
//
type CfnModuleDefaultVersionProps struct {
	// The Amazon Resource Name (ARN) of the module version to set as the default version.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduledefaultversion.html#cfn-cloudformation-moduledefaultversion-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name of the module.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduledefaultversion.html#cfn-cloudformation-moduledefaultversion-modulename
	//
	ModuleName *string `field:"optional" json:"moduleName" yaml:"moduleName"`
	// The ID for the specific version of the module.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduledefaultversion.html#cfn-cloudformation-moduledefaultversion-versionid
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

