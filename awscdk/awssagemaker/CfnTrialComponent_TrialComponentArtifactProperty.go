package awssagemaker


// Represents an input or output artifact of a trial component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trialComponentArtifactProperty := &TrialComponentArtifactProperty{
//   	Value: jsii.String("value"),
//
//   	// the properties below are optional
//   	MediaType: jsii.String("mediaType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-trialcomponent-trialcomponentartifact.html
//
type CfnTrialComponent_TrialComponentArtifactProperty struct {
	// The location of the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-trialcomponent-trialcomponentartifact.html#cfn-sagemaker-trialcomponent-trialcomponentartifact-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
	// The media type of the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-trialcomponent-trialcomponentartifact.html#cfn-sagemaker-trialcomponent-trialcomponentartifact-mediatype
	//
	MediaType *string `field:"optional" json:"mediaType" yaml:"mediaType"`
}

