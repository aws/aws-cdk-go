package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceCatalogProvisionedProductDetailsProperty := &serviceCatalogProvisionedProductDetailsProperty{
//   	provisionedProductId: jsii.String("provisionedProductId"),
//   	provisionedProductStatusMessage: jsii.String("provisionedProductStatusMessage"),
//   }
//
type CfnProject_ServiceCatalogProvisionedProductDetailsProperty struct {
	// `CfnProject.ServiceCatalogProvisionedProductDetailsProperty.ProvisionedProductId`.
	ProvisionedProductId *string `field:"optional" json:"provisionedProductId" yaml:"provisionedProductId"`
	// `CfnProject.ServiceCatalogProvisionedProductDetailsProperty.ProvisionedProductStatusMessage`.
	ProvisionedProductStatusMessage *string `field:"optional" json:"provisionedProductStatusMessage" yaml:"provisionedProductStatusMessage"`
}

