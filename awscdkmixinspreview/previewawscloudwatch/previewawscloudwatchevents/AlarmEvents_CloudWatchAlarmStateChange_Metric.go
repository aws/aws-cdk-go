package previewawscloudwatchevents


// Type definition for Metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metric := &Metric{
//   	Dimensions: []*string{
//   		jsii.String("dimensions"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Namespace: []*string{
//   		jsii.String("namespace"),
//   	},
//   }
//
// Experimental.
type AlarmEvents_CloudWatchAlarmStateChange_Metric struct {
	// dimensions property.
	//
	// Specify an array of string values to match this event if the actual value of dimensions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Dimensions *[]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// namespace property.
	//
	// Specify an array of string values to match this event if the actual value of namespace is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Namespace *[]*string `field:"optional" json:"namespace" yaml:"namespace"`
}

