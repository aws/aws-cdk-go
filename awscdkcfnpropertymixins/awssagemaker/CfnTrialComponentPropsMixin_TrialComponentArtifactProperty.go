package awssagemaker


// Represents an input or output artifact of a trial component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   trialComponentArtifactProperty := &TrialComponentArtifactProperty{
//   	MediaType: jsii.String("mediaType"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-trialcomponent-trialcomponentartifact.html
//
type CfnTrialComponentPropsMixin_TrialComponentArtifactProperty struct {
	// The media type of the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-trialcomponent-trialcomponentartifact.html#cfn-sagemaker-trialcomponent-trialcomponentartifact-mediatype
	//
	MediaType *string `field:"optional" json:"mediaType" yaml:"mediaType"`
	// The location of the artifact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-trialcomponent-trialcomponentartifact.html#cfn-sagemaker-trialcomponent-trialcomponentartifact-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

