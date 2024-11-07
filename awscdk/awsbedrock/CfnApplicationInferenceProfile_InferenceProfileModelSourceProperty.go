package awsbedrock


// Various ways to encode a list of models in a CreateInferenceProfile request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceProfileModelSourceProperty := &InferenceProfileModelSourceProperty{
//   	CopyFrom: jsii.String("copyFrom"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-applicationinferenceprofile-inferenceprofilemodelsource.html
//
type CfnApplicationInferenceProfile_InferenceProfileModelSourceProperty struct {
	// Source arns for a custom inference profile to copy its regional load balancing config from.
	//
	// This
	// can either be a foundation model or predefined inference profile ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-applicationinferenceprofile-inferenceprofilemodelsource.html#cfn-bedrock-applicationinferenceprofile-inferenceprofilemodelsource-copyfrom
	//
	CopyFrom *string `field:"required" json:"copyFrom" yaml:"copyFrom"`
}

