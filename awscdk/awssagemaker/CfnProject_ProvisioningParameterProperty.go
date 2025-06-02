package awssagemaker


// A key value pair used when you provision a project as a service catalog product.
//
// For information, see [What is AWS Service Catalog](https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisioningParameterProperty := &ProvisioningParameterProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-provisioningparameter.html
//
type CfnProject_ProvisioningParameterProperty struct {
	// The key that identifies a provisioning parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-provisioningparameter.html#cfn-sagemaker-project-provisioningparameter-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the provisioning parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-project-provisioningparameter.html#cfn-sagemaker-project-provisioningparameter-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

