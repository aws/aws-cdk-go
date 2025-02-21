package awsbedrock


// Contains information about the model or system-defined inference profile that is the source for an inference profile..
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
	// The ARN of the model or system-defined inference profile that is the source for the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-applicationinferenceprofile-inferenceprofilemodelsource.html#cfn-bedrock-applicationinferenceprofile-inferenceprofilemodelsource-copyfrom
	//
	CopyFrom *string `field:"required" json:"copyFrom" yaml:"copyFrom"`
}

