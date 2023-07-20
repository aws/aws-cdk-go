package awscdk


// Properties for defining a `CfnModuleVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModuleVersionProps := &CfnModuleVersionProps{
//   	ModuleName: jsii.String("moduleName"),
//   	ModulePackage: jsii.String("modulePackage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduleversion.html
//
type CfnModuleVersionProps struct {
	// The name of the module being registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduleversion.html#cfn-cloudformation-moduleversion-modulename
	//
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// A URL to the S3 bucket containing the package that contains the template fragment and schema files for the module version to register.
	//
	// > The user registering the module version must be able to access the module package in the S3 bucket. That's, the user needs to have [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) permissions for the package. For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduleversion.html#cfn-cloudformation-moduleversion-modulepackage
	//
	ModulePackage *string `field:"required" json:"modulePackage" yaml:"modulePackage"`
}

