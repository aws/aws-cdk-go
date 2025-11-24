package mixinsawscloudformation


// Properties for CfnModuleVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnModuleVersionMixinProps := &CfnModuleVersionMixinProps{
//   	ModuleName: jsii.String("moduleName"),
//   	ModulePackage: jsii.String("modulePackage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduleversion.html
//
type CfnModuleVersionMixinProps struct {
	// The name of the module being registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduleversion.html#cfn-cloudformation-moduleversion-modulename
	//
	ModuleName *string `field:"optional" json:"moduleName" yaml:"moduleName"`
	// A URL to the S3 bucket for the package that contains the template fragment and schema files for the module version to register.
	//
	// For more information, see [Module structure and requirements](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/modules-structure.html) in the *CloudFormation Command Line Interface (CLI) User Guide* .
	//
	// > To register the module version, you must have `s3:GetObject` permissions to access the S3 objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduleversion.html#cfn-cloudformation-moduleversion-modulepackage
	//
	ModulePackage *string `field:"optional" json:"modulePackage" yaml:"modulePackage"`
}

