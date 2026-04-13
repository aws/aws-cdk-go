package previewawselasticloadbalancingevents


// Type definition for HealthCheck_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   healthCheck1 := &HealthCheck1{
//   	HealthyThreshold: []*string{
//   		jsii.String("healthyThreshold"),
//   	},
//   	Interval: []*string{
//   		jsii.String("interval"),
//   	},
//   	Target: []*string{
//   		jsii.String("target"),
//   	},
//   	Timeout: []*string{
//   		jsii.String("timeout"),
//   	},
//   	UnhealthyThreshold: []*string{
//   		jsii.String("unhealthyThreshold"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_HealthCheck1 struct {
	// healthyThreshold property.
	//
	// Specify an array of string values to match this event if the actual value of healthyThreshold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HealthyThreshold *[]*string `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// interval property.
	//
	// Specify an array of string values to match this event if the actual value of interval is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Interval *[]*string `field:"optional" json:"interval" yaml:"interval"`
	// target property.
	//
	// Specify an array of string values to match this event if the actual value of target is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Target *[]*string `field:"optional" json:"target" yaml:"target"`
	// timeout property.
	//
	// Specify an array of string values to match this event if the actual value of timeout is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Timeout *[]*string `field:"optional" json:"timeout" yaml:"timeout"`
	// unhealthyThreshold property.
	//
	// Specify an array of string values to match this event if the actual value of unhealthyThreshold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnhealthyThreshold *[]*string `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

