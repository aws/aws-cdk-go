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
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	Buildspec: []*string{
//   		jsii.String("buildspec"),
//   	},
//   	InsecureSsl: []*string{
//   		jsii.String("insecureSsl"),
//   	},
//   	Location: []*string{
//   		jsii.String("location"),
//   	},
//   	ReportBuildStatus: []*string{
//   		jsii.String("reportBuildStatus"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Source struct {
	// auth property.
	//
	// Specify an array of string values to match this event if the actual value of auth is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Auth *AWSAPICallViaCloudTrail_Auth `field:"optional" json:"auth" yaml:"auth"`
	// buildspec property.
	//
	// Specify an array of string values to match this event if the actual value of buildspec is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Buildspec *[]*string `field:"optional" json:"buildspec" yaml:"buildspec"`
	// insecureSsl property.
	//
	// Specify an array of string values to match this event if the actual value of insecureSsl is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InsecureSsl *[]*string `field:"optional" json:"insecureSsl" yaml:"insecureSsl"`
	// location property.
	//
	// Specify an array of string values to match this event if the actual value of location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Location *[]*string `field:"optional" json:"location" yaml:"location"`
	// reportBuildStatus property.
	//
	// Specify an array of string values to match this event if the actual value of reportBuildStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReportBuildStatus *[]*string `field:"optional" json:"reportBuildStatus" yaml:"reportBuildStatus"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

