package previewawscodebuildevents


// Type definition for Artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   artifact := &Artifact{
//   	Location: []*string{
//   		jsii.String("location"),
//   	},
//   	Md5Sum: []*string{
//   		jsii.String("md5Sum"),
//   	},
//   	Sha256Sum: []*string{
//   		jsii.String("sha256Sum"),
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildStateChange_Artifact struct {
	// location property.
	//
	// Specify an array of string values to match this event if the actual value of location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Location *[]*string `field:"optional" json:"location" yaml:"location"`
	// md5sum property.
	//
	// Specify an array of string values to match this event if the actual value of md5sum is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Md5Sum *[]*string `field:"optional" json:"md5Sum" yaml:"md5Sum"`
	// sha256sum property.
	//
	// Specify an array of string values to match this event if the actual value of sha256sum is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Sha256Sum *[]*string `field:"optional" json:"sha256Sum" yaml:"sha256Sum"`
}

