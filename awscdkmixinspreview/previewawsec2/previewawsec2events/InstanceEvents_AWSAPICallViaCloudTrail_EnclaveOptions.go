package previewawsec2events


// Type definition for EnclaveOptions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   enclaveOptions := &EnclaveOptions{
//   	Enabled: []*string{
//   		jsii.String("enabled"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_EnclaveOptions struct {
	// enabled property.
	//
	// Specify an array of string values to match this event if the actual value of enabled is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Enabled *[]*string `field:"optional" json:"enabled" yaml:"enabled"`
}

