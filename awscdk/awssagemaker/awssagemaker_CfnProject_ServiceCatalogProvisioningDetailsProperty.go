package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceCatalogProvisioningDetailsProperty := &serviceCatalogProvisioningDetailsProperty{
//   	productId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	pathId: jsii.String("pathId"),
//   	provisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	provisioningParameters: []interface{}{
//   		&provisioningParameterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProject_ServiceCatalogProvisioningDetailsProperty struct {
	// `CfnProject.ServiceCatalogProvisioningDetailsProperty.ProductId`.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// `CfnProject.ServiceCatalogProvisioningDetailsProperty.PathId`.
	PathId *string `field:"optional" json:"pathId" yaml:"pathId"`
	// `CfnProject.ServiceCatalogProvisioningDetailsProperty.ProvisioningArtifactId`.
	ProvisioningArtifactId *string `field:"optional" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// `CfnProject.ServiceCatalogProvisioningDetailsProperty.ProvisioningParameters`.
	ProvisioningParameters interface{} `field:"optional" json:"provisioningParameters" yaml:"provisioningParameters"`
}

