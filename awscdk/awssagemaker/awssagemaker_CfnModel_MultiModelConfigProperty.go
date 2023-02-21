package awssagemaker


// Specifies additional configuration for hosting multi-model endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiModelConfigProperty := &multiModelConfigProperty{
//   	modelCacheSetting: jsii.String("modelCacheSetting"),
//   }
//
type CfnModel_MultiModelConfigProperty struct {
	// Whether to cache models for a multi-model endpoint.
	//
	// By default, multi-model endpoints cache models so that a model does not have to be loaded into memory each time it is invoked. Some use cases do not benefit from model caching. For example, if an endpoint hosts a large number of models that are each invoked infrequently, the endpoint might perform better if you disable model caching. To disable model caching, set the value of this parameter to Disabled.
	ModelCacheSetting *string `field:"optional" json:"modelCacheSetting" yaml:"modelCacheSetting"`
}

