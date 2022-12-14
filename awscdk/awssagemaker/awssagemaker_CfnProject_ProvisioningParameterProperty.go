package awssagemaker


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
type CfnProject_ProvisioningParameterProperty struct {
	// `CfnProject.ProvisioningParameterProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnProject.ProvisioningParameterProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

