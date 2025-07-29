package awsappflow


// The properties that are applied when Google Analytics is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   googleAnalyticsSourcePropertiesProperty := &GoogleAnalyticsSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-googleanalyticssourceproperties.html
//
type CfnFlow_GoogleAnalyticsSourcePropertiesProperty struct {
	// The object specified in the Google Analytics flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-googleanalyticssourceproperties.html#cfn-appflow-flow-googleanalyticssourceproperties-object
	//
	Object *string `field:"required" json:"object" yaml:"object"`
}

