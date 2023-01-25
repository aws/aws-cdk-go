package awssagemaker


// Details that you specify to provision a service catalog product.
//
// For information about service catalog, see [What is AWS Service Catalog](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html) .
//
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
	// The ID of the product to provision.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// The path identifier of the product.
	//
	// This value is optional if the product has a default path, and required if the product has more than one path.
	PathId *string `field:"optional" json:"pathId" yaml:"pathId"`
	// The ID of the provisioning artifact.
	ProvisioningArtifactId *string `field:"optional" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// A list of key value pairs that you specify when you provision a product.
	ProvisioningParameters interface{} `field:"optional" json:"provisioningParameters" yaml:"provisioningParameters"`
}

