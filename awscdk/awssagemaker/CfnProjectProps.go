package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var serviceCatalogProvisioningDetails interface{}
//
//   cfnProjectProps := &CfnProjectProps{
//   	ProjectName: jsii.String("projectName"),
//
//   	// the properties below are optional
//   	ProjectDescription: jsii.String("projectDescription"),
//   	ServiceCatalogProvisionedProductDetails: &ServiceCatalogProvisionedProductDetailsProperty{
//   		ProvisionedProductId: jsii.String("provisionedProductId"),
//   		ProvisionedProductStatusMessage: jsii.String("provisionedProductStatusMessage"),
//   	},
//   	ServiceCatalogProvisioningDetails: serviceCatalogProvisioningDetails,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateProviderDetails: []interface{}{
//   		&TemplateProviderDetailProperty{
//   			CfnTemplateProviderDetail: &CfnTemplateProviderDetailProperty{
//   				TemplateName: jsii.String("templateName"),
//   				TemplateUrl: jsii.String("templateUrl"),
//
//   				// the properties below are optional
//   				Parameters: []interface{}{
//   					&CfnStackParameterProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-project.html
//
type CfnProjectProps struct {
	// The name of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-project.html#cfn-sagemaker-project-projectname
	//
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// The description of the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-project.html#cfn-sagemaker-project-projectdescription
	//
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// Details of a provisioned service catalog product.
	//
	// For information about service catalog, see [What is AWS Service Catalog](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-project.html#cfn-sagemaker-project-servicecatalogprovisionedproductdetails
	//
	ServiceCatalogProvisionedProductDetails interface{} `field:"optional" json:"serviceCatalogProvisionedProductDetails" yaml:"serviceCatalogProvisionedProductDetails"`
	// The product ID and provisioning artifact ID to provision a service catalog.
	//
	// For information, see [What is AWS Service Catalog](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-project.html#cfn-sagemaker-project-servicecatalogprovisioningdetails
	//
	ServiceCatalogProvisioningDetails interface{} `field:"optional" json:"serviceCatalogProvisioningDetails" yaml:"serviceCatalogProvisioningDetails"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-project.html#cfn-sagemaker-project-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An array of template providers associated with the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-project.html#cfn-sagemaker-project-templateproviderdetails
	//
	TemplateProviderDetails interface{} `field:"optional" json:"templateProviderDetails" yaml:"templateProviderDetails"`
}

