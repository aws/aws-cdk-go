package awsservicecatalog


// Information about a provisioning artifact (also known as a version) for a product.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var info interface{}
//
//   provisioningArtifactPropertiesProperty := &ProvisioningArtifactPropertiesProperty{
//   	Info: info,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DisableTemplateValidation: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   }
//
type CfnCloudFormationProduct_ProvisioningArtifactPropertiesProperty struct {
	// Specify the template source with one of the following options, but not both.
	//
	// Keys accepted: [ `LoadTemplateFromURL` , `ImportFromPhysicalId` ]
	//
	// The URL of the AWS CloudFormation template in Amazon S3 in JSON format. Specify the URL in JSON format as follows:
	//
	// `"LoadTemplateFromURL": "https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/..."`
	//
	// `ImportFromPhysicalId` : The physical id of the resource that contains the template. Currently only supports AWS CloudFormation stack arn. Specify the physical id in JSON format as follows: `ImportFromPhysicalId: “arn:aws:cloudformation:[us-east-1]:[accountId]:stack/[StackName]/[resourceId]`
	Info interface{} `field:"required" json:"info" yaml:"info"`
	// The description of the provisioning artifact, including how it differs from the previous provisioning artifact.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set to true, AWS Service Catalog stops validating the specified provisioning artifact even if it is invalid.
	DisableTemplateValidation interface{} `field:"optional" json:"disableTemplateValidation" yaml:"disableTemplateValidation"`
	// The name of the provisioning artifact (for example, v1 v2beta).
	//
	// No spaces are allowed.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

