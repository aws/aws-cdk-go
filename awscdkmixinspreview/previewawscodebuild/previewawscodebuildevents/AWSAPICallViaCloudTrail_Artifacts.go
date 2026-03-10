package previewawscodebuildevents


// Type definition for Artifacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   artifacts := &Artifacts{
//   	EncryptionDisabled: []*string{
//   		jsii.String("encryptionDisabled"),
//   	},
//   	Location: []*string{
//   		jsii.String("location"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Artifacts struct {
	// encryptionDisabled property.
	//
	// Specify an array of string values to match this event if the actual value of encryptionDisabled is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EncryptionDisabled *[]*string `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// location property.
	//
	// Specify an array of string values to match this event if the actual value of location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Location *[]*string `field:"optional" json:"location" yaml:"location"`
}

