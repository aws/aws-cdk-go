package awscustomerprofiles


// The properties that are applied when ServiceNow is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowSourcePropertiesProperty := &serviceNowSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnIntegration_ServiceNowSourcePropertiesProperty struct {
	// The object specified in the ServiceNow flow source.
	Object *string `field:"required" json:"object" yaml:"object"`
}

