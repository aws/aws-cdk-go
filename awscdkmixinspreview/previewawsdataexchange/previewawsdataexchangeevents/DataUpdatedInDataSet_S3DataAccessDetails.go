package previewawsdataexchangeevents


// Type definition for S3DataAccessDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3DataAccessDetails := &S3DataAccessDetails{
//   	KeyPrefixes: []*string{
//   		jsii.String("keyPrefixes"),
//   	},
//   	Keys: []*string{
//   		jsii.String("keys"),
//   	},
//   }
//
// Experimental.
type DataUpdatedInDataSet_S3DataAccessDetails struct {
	// KeyPrefixes property.
	//
	// Specify an array of string values to match this event if the actual value of KeyPrefixes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KeyPrefixes *[]*string `field:"optional" json:"keyPrefixes" yaml:"keyPrefixes"`
	// Keys property.
	//
	// Specify an array of string values to match this event if the actual value of Keys is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
}

