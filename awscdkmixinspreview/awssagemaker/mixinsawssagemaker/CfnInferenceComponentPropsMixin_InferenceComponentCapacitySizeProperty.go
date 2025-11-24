package mixinsawssagemaker


// Specifies the type and size of the endpoint capacity to activate for a rolling deployment or a rollback strategy.
//
// You can specify your batches as either of the following:
//
// - A count of inference component copies
// - The overall percentage or your fleet
//
// For a rollback strategy, if you don't specify the fields in this object, or if you set the `Value` parameter to 100%, then SageMaker AI uses a blue/green rollback strategy and rolls all traffic back to the blue fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inferenceComponentCapacitySizeProperty := &InferenceComponentCapacitySizeProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcapacitysize.html
//
type CfnInferenceComponentPropsMixin_InferenceComponentCapacitySizeProperty struct {
	// Specifies the endpoint capacity type.
	//
	// - **COPY_COUNT** - The endpoint activates based on the number of inference component copies.
	// - **CAPACITY_PERCENT** - The endpoint activates based on the specified percentage of capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcapacitysize.html#cfn-sagemaker-inferencecomponent-inferencecomponentcapacitysize-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Defines the capacity size, either as a number of inference component copies or a capacity percentage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcapacitysize.html#cfn-sagemaker-inferencecomponent-inferencecomponentcapacitysize-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

