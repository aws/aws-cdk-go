package previewawss3events


// Type definition for TlsDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tlsDetails := &TlsDetails{
//   	CipherSuite: []*string{
//   		jsii.String("cipherSuite"),
//   	},
//   	ClientProvidedHostHeader: []*string{
//   		jsii.String("clientProvidedHostHeader"),
//   	},
//   	TlsVersion: []*string{
//   		jsii.String("tlsVersion"),
//   	},
//   }
//
// Experimental.
type BucketEvents_AWSAPICallViaCloudTrail_TlsDetails struct {
	// cipherSuite property.
	//
	// Specify an array of string values to match this event if the actual value of cipherSuite is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CipherSuite *[]*string `field:"optional" json:"cipherSuite" yaml:"cipherSuite"`
	// clientProvidedHostHeader property.
	//
	// Specify an array of string values to match this event if the actual value of clientProvidedHostHeader is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientProvidedHostHeader *[]*string `field:"optional" json:"clientProvidedHostHeader" yaml:"clientProvidedHostHeader"`
	// tlsVersion property.
	//
	// Specify an array of string values to match this event if the actual value of tlsVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TlsVersion *[]*string `field:"optional" json:"tlsVersion" yaml:"tlsVersion"`
}

