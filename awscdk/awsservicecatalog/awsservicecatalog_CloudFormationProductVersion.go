package awsservicecatalog


// Properties of product version (also known as a provisioning artifact).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudFormationTemplate cloudFormationTemplate
//
//   cloudFormationProductVersion := &CloudFormationProductVersion{
//   	CloudFormationTemplate: cloudFormationTemplate,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ProductVersionName: jsii.String("productVersionName"),
//   	ValidateTemplate: jsii.Boolean(false),
//   }
//
// Experimental.
type CloudFormationProductVersion struct {
	// The S3 template that points to the provisioning version template.
	// Experimental.
	CloudFormationTemplate CloudFormationTemplate `field:"required" json:"cloudFormationTemplate" yaml:"cloudFormationTemplate"`
	// The description of the product version.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the product version.
	// Experimental.
	ProductVersionName *string `field:"optional" json:"productVersionName" yaml:"productVersionName"`
	// Whether the specified product template will be validated by CloudFormation.
	//
	// If turned off, an invalid template configuration can be stored.
	// Experimental.
	ValidateTemplate *bool `field:"optional" json:"validateTemplate" yaml:"validateTemplate"`
}

