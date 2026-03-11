package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   provisionedConcurrencyConfigProperty := &ProvisionedConcurrencyConfigProperty{
//   	ProvisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-provisionedconcurrencyconfig.html
//
type CfnFunctionPropsMixin_ProvisionedConcurrencyConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-provisionedconcurrencyconfig.html#cfn-serverless-function-provisionedconcurrencyconfig-provisionedconcurrentexecutions
	//
	ProvisionedConcurrentExecutions *string `field:"optional" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

