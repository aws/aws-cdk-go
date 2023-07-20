package awssagemaker


// Details of a provisioned service catalog product.
//
// For information about service catalog, see [What is AWS Service Catalog](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceCatalogProvisionedProductDetailsProperty := &ServiceCatalogProvisionedProductDetailsProperty{
//   	ProvisionedProductId: jsii.String("provisionedProductId"),
//   	ProvisionedProductStatusMessage: jsii.String("provisionedProductStatusMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisionedproductdetails.html
//
type CfnProject_ServiceCatalogProvisionedProductDetailsProperty struct {
	// The ID of the provisioned product.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisionedproductdetails.html#cfn-sagemaker-project-servicecatalogprovisionedproductdetails-provisionedproductid
	//
	ProvisionedProductId *string `field:"optional" json:"provisionedProductId" yaml:"provisionedProductId"`
	// The current status of the product.
	//
	// - `AVAILABLE` - Stable state, ready to perform any operation. The most recent operation succeeded and completed.
	// - `UNDER_CHANGE` - Transitive state. Operations performed might not have valid results. Wait for an AVAILABLE status before performing operations.
	// - `TAINTED` - Stable state, ready to perform any operation. The stack has completed the requested operation but is not exactly what was requested. For example, a request to update to a new version failed and the stack rolled back to the current version.
	// - `ERROR` - An unexpected error occurred. The provisioned product exists but the stack is not running. For example, CloudFormation received a parameter value that was not valid and could not launch the stack.
	// - `PLAN_IN_PROGRESS` - Transitive state. The plan operations were performed to provision a new product, but resources have not yet been created. After reviewing the list of resources to be created, execute the plan. Wait for an AVAILABLE status before performing operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-servicecatalogprovisionedproductdetails.html#cfn-sagemaker-project-servicecatalogprovisionedproductdetails-provisionedproductstatusmessage
	//
	ProvisionedProductStatusMessage *string `field:"optional" json:"provisionedProductStatusMessage" yaml:"provisionedProductStatusMessage"`
}

