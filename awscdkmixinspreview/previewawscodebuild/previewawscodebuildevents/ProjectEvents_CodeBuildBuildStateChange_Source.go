package previewawscodebuildevents


// Type definition for Source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   source := &Source{
//   	Auth: &Auth{
//   		Resource: []*string{
//   			jsii.String("resource"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	Buildspec: []*string{
//   		jsii.String("buildspec"),
//   	},
//   	Location: []*string{
//   		jsii.String("location"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type ProjectEvents_CodeBuildBuildStateChange_Source struct {
	// auth property.
	//
	// Specify an array of string values to match this event if the actual value of auth is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Auth *ProjectEvents_CodeBuildBuildStateChange_Auth `field:"optional" json:"auth" yaml:"auth"`
	// buildspec property.
	//
	// Specify an array of string values to match this event if the actual value of buildspec is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Buildspec *[]*string `field:"optional" json:"buildspec" yaml:"buildspec"`
	// location property.
	//
	// Specify an array of string values to match this event if the actual value of location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Location *[]*string `field:"optional" json:"location" yaml:"location"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

