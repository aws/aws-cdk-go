package previewawswellarchitectedevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	LensArn: []*string{
//   		jsii.String("lensArn"),
//   	},
//   	LensVersion: []*string{
//   		jsii.String("lensVersion"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// LensArn property.
	//
	// Specify an array of string values to match this event if the actual value of LensArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LensArn *[]*string `field:"optional" json:"lensArn" yaml:"lensArn"`
	// LensVersion property.
	//
	// Specify an array of string values to match this event if the actual value of LensVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LensVersion *[]*string `field:"optional" json:"lensVersion" yaml:"lensVersion"`
}

