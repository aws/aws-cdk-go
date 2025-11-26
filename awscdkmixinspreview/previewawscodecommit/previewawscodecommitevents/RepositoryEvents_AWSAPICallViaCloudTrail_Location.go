package previewawscodecommitevents


// Type definition for Location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   location := &Location{
//   	FilePath: []*string{
//   		jsii.String("filePath"),
//   	},
//   	FilePosition: []*string{
//   		jsii.String("filePosition"),
//   	},
//   	RelativeFileVersion: []*string{
//   		jsii.String("relativeFileVersion"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_Location struct {
	// filePath property.
	//
	// Specify an array of string values to match this event if the actual value of filePath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilePath *[]*string `field:"optional" json:"filePath" yaml:"filePath"`
	// filePosition property.
	//
	// Specify an array of string values to match this event if the actual value of filePosition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FilePosition *[]*string `field:"optional" json:"filePosition" yaml:"filePosition"`
	// relativeFileVersion property.
	//
	// Specify an array of string values to match this event if the actual value of relativeFileVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RelativeFileVersion *[]*string `field:"optional" json:"relativeFileVersion" yaml:"relativeFileVersion"`
}

