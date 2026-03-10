package previewawsiotanalyticsevents


// Type definition for RetentionPeriod.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retentionPeriod := &RetentionPeriod{
//   	NumberOfDays: []*string{
//   		jsii.String("numberOfDays"),
//   	},
//   	Unlimited: []*string{
//   		jsii.String("unlimited"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RetentionPeriod struct {
	// numberOfDays property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfDays is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfDays *[]*string `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// unlimited property.
	//
	// Specify an array of string values to match this event if the actual value of unlimited is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Unlimited *[]*string `field:"optional" json:"unlimited" yaml:"unlimited"`
}

