package previewawscodecommitevents


// Type definition for ResponseElementsItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElementsItem := &ResponseElementsItem{
//   	AbsolutePath: []*string{
//   		jsii.String("absolutePath"),
//   	},
//   	BlobId: []*string{
//   		jsii.String("blobId"),
//   	},
//   	FileMode: []*string{
//   		jsii.String("fileMode"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElementsItem struct {
	// absolutePath property.
	//
	// Specify an array of string values to match this event if the actual value of absolutePath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AbsolutePath *[]*string `field:"optional" json:"absolutePath" yaml:"absolutePath"`
	// blobId property.
	//
	// Specify an array of string values to match this event if the actual value of blobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BlobId *[]*string `field:"optional" json:"blobId" yaml:"blobId"`
	// fileMode property.
	//
	// Specify an array of string values to match this event if the actual value of fileMode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FileMode *[]*string `field:"optional" json:"fileMode" yaml:"fileMode"`
}

