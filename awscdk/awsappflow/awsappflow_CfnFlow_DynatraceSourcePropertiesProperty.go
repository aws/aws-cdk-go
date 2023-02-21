package awsappflow


// The properties that are applied when Dynatrace is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynatraceSourcePropertiesProperty := &dynatraceSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_DynatraceSourcePropertiesProperty struct {
	// The object specified in the Dynatrace flow source.
	Object *string `field:"required" json:"object" yaml:"object"`
}

