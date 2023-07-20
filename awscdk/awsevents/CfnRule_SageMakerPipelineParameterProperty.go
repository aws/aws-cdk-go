package awsevents


// Name/Value pair of a parameter to start execution of a SageMaker Model Building Pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerPipelineParameterProperty := &SageMakerPipelineParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-sagemakerpipelineparameter.html
//
type CfnRule_SageMakerPipelineParameterProperty struct {
	// Name of parameter to start execution of a SageMaker Model Building Pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-sagemakerpipelineparameter.html#cfn-events-rule-sagemakerpipelineparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of parameter to start execution of a SageMaker Model Building Pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-sagemakerpipelineparameter.html#cfn-events-rule-sagemakerpipelineparameter-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

