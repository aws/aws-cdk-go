package previewawscodebuildevents


// Type definition for Auth.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   auth := &Auth{
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildPhaseChange_Auth struct {
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

