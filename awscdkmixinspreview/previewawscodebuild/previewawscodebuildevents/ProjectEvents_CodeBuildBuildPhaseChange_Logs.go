package previewawscodebuildevents


// Type definition for Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logs := &Logs{
//   	DeepLink: []*string{
//   		jsii.String("deepLink"),
//   	},
//   	GroupName: []*string{
//   		jsii.String("groupName"),
//   	},
//   	StreamName: []*string{
//   		jsii.String("streamName"),
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildPhaseChange_Logs struct {
	// deep-link property.
	//
	// Specify an array of string values to match this event if the actual value of deep-link is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeepLink *[]*string `field:"optional" json:"deepLink" yaml:"deepLink"`
	// group-name property.
	//
	// Specify an array of string values to match this event if the actual value of group-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GroupName *[]*string `field:"optional" json:"groupName" yaml:"groupName"`
	// stream-name property.
	//
	// Specify an array of string values to match this event if the actual value of stream-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StreamName *[]*string `field:"optional" json:"streamName" yaml:"streamName"`
}

