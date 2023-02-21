package awscloudformation


// Properties for defining a `CfnModuleDefaultVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModuleDefaultVersionProps := &cfnModuleDefaultVersionProps{
//   	arn: jsii.String("arn"),
//   	moduleName: jsii.String("moduleName"),
//   	versionId: jsii.String("versionId"),
//   }
//
type CfnModuleDefaultVersionProps struct {
	// The Amazon Resource Name (ARN) of the module version to set as the default version.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name of the module.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	ModuleName *string `field:"optional" json:"moduleName" yaml:"moduleName"`
	// The ID for the specific version of the module.
	//
	// Conditional: You must specify either `Arn` , or `ModuleName` and `VersionId` .
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

