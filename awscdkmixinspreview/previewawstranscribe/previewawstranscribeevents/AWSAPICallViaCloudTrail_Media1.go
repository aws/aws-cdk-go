package previewawstranscribeevents


// Type definition for Media_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   media1 := &Media1{
//   	MediaFileUri: []*string{
//   		jsii.String("mediaFileUri"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Media1 struct {
	// mediaFileUri property.
	//
	// Specify an array of string values to match this event if the actual value of mediaFileUri is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MediaFileUri *[]*string `field:"optional" json:"mediaFileUri" yaml:"mediaFileUri"`
}

