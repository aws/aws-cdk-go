package previewawscodebuildevents


// Type definition for GitSubmodulesConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gitSubmodulesConfig := &GitSubmodulesConfig{
//   	FetchSubmodules: []*string{
//   		jsii.String("fetchSubmodules"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_GitSubmodulesConfig struct {
	// fetchSubmodules property.
	//
	// Specify an array of string values to match this event if the actual value of fetchSubmodules is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FetchSubmodules *[]*string `field:"optional" json:"fetchSubmodules" yaml:"fetchSubmodules"`
}

