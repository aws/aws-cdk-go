package awssagemaker


// Settings that affect how the inference component caches data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceComponentDataCacheConfigProperty := &InferenceComponentDataCacheConfigProperty{
//   	EnableCaching: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentdatacacheconfig.html
//
type CfnInferenceComponentPropsMixin_InferenceComponentDataCacheConfigProperty struct {
	// Whether the endpoint caches the model artifacts and container image on each instance it provisions for the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentdatacacheconfig.html#cfn-sagemaker-inferencecomponent-inferencecomponentdatacacheconfig-enablecaching
	//
	EnableCaching interface{} `field:"optional" json:"enableCaching" yaml:"enableCaching"`
}

