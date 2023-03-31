package awscloudformation


// Properties for defining a `CfnModuleVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModuleVersionProps := &cfnModuleVersionProps{
//   	moduleName: jsii.String("moduleName"),
//   	modulePackage: jsii.String("modulePackage"),
//   }
//
type CfnModuleVersionProps struct {
	// The name of the module being registered.
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// A URL to the S3 bucket containing the package that contains the template fragment and schema files for the module version to register.
	//
	// > The user registering the module version must be able to access the module package in the S3 bucket. That's, the user needs to have [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) permissions for the package. For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	ModulePackage *string `field:"required" json:"modulePackage" yaml:"modulePackage"`
}

