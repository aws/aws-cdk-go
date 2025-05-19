package awsappflow


// The properties that are applied when Datadog is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datadogSourcePropertiesProperty := &DatadogSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-datadogsourceproperties.html
//
type CfnFlow_DatadogSourcePropertiesProperty struct {
	// The object specified in the Datadog flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-datadogsourceproperties.html#cfn-appflow-flow-datadogsourceproperties-object
	//
	Object *string `field:"required" json:"object" yaml:"object"`
}

