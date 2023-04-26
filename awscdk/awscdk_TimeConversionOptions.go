// An experiment to bundle the entire CDK into a single module
package awscdk


// Options for how to convert time to a different unit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   timeConversionOptions := &TimeConversionOptions{
//   	Integral: jsii.Boolean(false),
//   }
//
// Experimental.
type TimeConversionOptions struct {
	// If `true`, conversions into a larger time unit (e.g. `Seconds` to `Minutes`) will fail if the result is not an integer.
	// Experimental.
	Integral *bool `field:"optional" json:"integral" yaml:"integral"`
}

