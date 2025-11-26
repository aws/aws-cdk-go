package previewawsecrevents


// Type definition for RequestParametersItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParametersItem := &RequestParametersItem{
//   	ImageTag: []*string{
//   		jsii.String("imageTag"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem struct {
	// imageTag property.
	//
	// Specify an array of string values to match this event if the actual value of imageTag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageTag *[]*string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

