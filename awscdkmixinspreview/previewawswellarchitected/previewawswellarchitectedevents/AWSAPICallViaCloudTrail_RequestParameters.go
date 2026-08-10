package previewawswellarchitectedevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	ClientRequestToken: []*string{
//   		jsii.String("clientRequestToken"),
//   	},
//   	IsMajorVersion: []*string{
//   		jsii.String("isMajorVersion"),
//   	},
//   	LensAlias: []*string{
//   		jsii.String("lensAlias"),
//   	},
//   	LensVersion: []*string{
//   		jsii.String("lensVersion"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParameters struct {
	// ClientRequestToken property.
	//
	// Specify an array of string values to match this event if the actual value of ClientRequestToken is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientRequestToken *[]*string `field:"optional" json:"clientRequestToken" yaml:"clientRequestToken"`
	// IsMajorVersion property.
	//
	// Specify an array of string values to match this event if the actual value of IsMajorVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IsMajorVersion *[]*string `field:"optional" json:"isMajorVersion" yaml:"isMajorVersion"`
	// LensAlias property.
	//
	// Specify an array of string values to match this event if the actual value of LensAlias is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LensAlias *[]*string `field:"optional" json:"lensAlias" yaml:"lensAlias"`
	// LensVersion property.
	//
	// Specify an array of string values to match this event if the actual value of LensVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LensVersion *[]*string `field:"optional" json:"lensVersion" yaml:"lensVersion"`
}

