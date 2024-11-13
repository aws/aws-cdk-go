package awsbedrock


// Contains information about a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceProfileModelProperty := &InferenceProfileModelProperty{
//   	ModelArn: jsii.String("modelArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-applicationinferenceprofile-inferenceprofilemodel.html
//
type CfnApplicationInferenceProfile_InferenceProfileModelProperty struct {
	// The Amazon Resource Name (ARN) of the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-applicationinferenceprofile-inferenceprofilemodel.html#cfn-bedrock-applicationinferenceprofile-inferenceprofilemodel-modelarn
	//
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
}

