package previewawsappflowmixins


// The properties that are applied when Amplitude is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   amplitudeSourcePropertiesProperty := &AmplitudeSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-amplitudesourceproperties.html
//
type CfnFlowPropsMixin_AmplitudeSourcePropertiesProperty struct {
	// The object specified in the Amplitude flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-amplitudesourceproperties.html#cfn-appflow-flow-amplitudesourceproperties-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

