package awsobservabilityadmin


// Structure containing a name field limited to 64 characters for header or query parameter identification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleHeaderProperty := &SingleHeaderProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-singleheader.html
//
type CfnTelemetryRule_SingleHeaderProperty struct {
	// The name value, limited to 64 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-singleheader.html#cfn-observabilityadmin-telemetryrule-singleheader-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

