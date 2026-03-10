package previewawsimagebuilderevents


// Type definition for Finding-severity-counts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   findingSeverityCounts := &FindingSeverityCounts{
//   	All: []*string{
//   		jsii.String("all"),
//   	},
//   	Critical: []*string{
//   		jsii.String("critical"),
//   	},
//   	High: []*string{
//   		jsii.String("high"),
//   	},
//   	Medium: []*string{
//   		jsii.String("medium"),
//   	},
//   }
//
// Experimental.
type EC2ImageBuilderCVEDetected_FindingSeverityCounts struct {
	// all property.
	//
	// Specify an array of string values to match this event if the actual value of all is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	All *[]*string `field:"optional" json:"all" yaml:"all"`
	// critical property.
	//
	// Specify an array of string values to match this event if the actual value of critical is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Critical *[]*string `field:"optional" json:"critical" yaml:"critical"`
	// high property.
	//
	// Specify an array of string values to match this event if the actual value of high is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	High *[]*string `field:"optional" json:"high" yaml:"high"`
	// medium property.
	//
	// Specify an array of string values to match this event if the actual value of medium is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Medium *[]*string `field:"optional" json:"medium" yaml:"medium"`
}

