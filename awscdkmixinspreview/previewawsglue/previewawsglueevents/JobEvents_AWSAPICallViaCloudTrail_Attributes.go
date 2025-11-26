package previewawsglueevents


// Type definition for Attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attributes := &Attributes{
//   	CreationDate: []*string{
//   		jsii.String("creationDate"),
//   	},
//   	MfaAuthenticated: []*string{
//   		jsii.String("mfaAuthenticated"),
//   	},
//   }
//
// Experimental.
type JobEvents_AWSAPICallViaCloudTrail_Attributes struct {
	// creationDate property.
	//
	// Specify an array of string values to match this event if the actual value of creationDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationDate *[]*string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// mfaAuthenticated property.
	//
	// Specify an array of string values to match this event if the actual value of mfaAuthenticated is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MfaAuthenticated *[]*string `field:"optional" json:"mfaAuthenticated" yaml:"mfaAuthenticated"`
}

