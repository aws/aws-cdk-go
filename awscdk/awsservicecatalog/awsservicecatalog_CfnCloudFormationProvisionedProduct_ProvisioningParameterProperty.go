package awsservicecatalog


// Information about a parameter used to provision a product.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisioningParameterProperty := &provisioningParameterProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnCloudFormationProvisionedProduct_ProvisioningParameterProperty struct {
	// The parameter key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The parameter value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

