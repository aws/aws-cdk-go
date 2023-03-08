package awsappflow


// The properties that are applied when Amplitude is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amplitudeSourcePropertiesProperty := &amplitudeSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_AmplitudeSourcePropertiesProperty struct {
	// The object specified in the Amplitude flow source.
	Object *string `field:"required" json:"object" yaml:"object"`
}

