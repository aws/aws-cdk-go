package previewawscodecommitevents


// Type definition for RequestParametersItem_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParametersItem1 := &RequestParametersItem1{
//   	Commit: []*string{
//   		jsii.String("commit"),
//   	},
//   	Ref: []*string{
//   		jsii.String("ref"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem1 struct {
	// commit property.
	//
	// Specify an array of string values to match this event if the actual value of commit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Commit *[]*string `field:"optional" json:"commit" yaml:"commit"`
	// ref property.
	//
	// Specify an array of string values to match this event if the actual value of ref is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ref *[]*string `field:"optional" json:"ref" yaml:"ref"`
}

