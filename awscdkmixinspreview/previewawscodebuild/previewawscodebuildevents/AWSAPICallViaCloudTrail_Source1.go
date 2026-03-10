package previewawscodebuildevents


// Type definition for Source_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   source1 := &Source1{
//   	Auth: &Auth{
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	Buildspec: []*string{
//   		jsii.String("buildspec"),
//   	},
//   	GitCloneDepth: []*string{
//   		jsii.String("gitCloneDepth"),
//   	},
//   	GitSubmodulesConfig: &GitSubmodulesConfig{
//   		FetchSubmodules: []*string{
//   			jsii.String("fetchSubmodules"),
//   		},
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
type AWSAPICallViaCloudTrail_Source1 struct {
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
	// gitCloneDepth property.
	//
	// Specify an array of string values to match this event if the actual value of gitCloneDepth is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GitCloneDepth *[]*string `field:"optional" json:"gitCloneDepth" yaml:"gitCloneDepth"`
	// gitSubmodulesConfig property.
	//
	// Specify an array of string values to match this event if the actual value of gitSubmodulesConfig is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GitSubmodulesConfig *AWSAPICallViaCloudTrail_GitSubmodulesConfig `field:"optional" json:"gitSubmodulesConfig" yaml:"gitSubmodulesConfig"`
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

