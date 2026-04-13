package previewawselasticloadbalancingevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	DNsName: []*string{
//   		jsii.String("dNsName"),
//   	},
//   	HealthCheck: &HealthCheck{
//   		HealthyThreshold: []*string{
//   			jsii.String("healthyThreshold"),
//   		},
//   		Interval: []*string{
//   			jsii.String("interval"),
//   		},
//   		Target: []*string{
//   			jsii.String("target"),
//   		},
//   		Timeout: []*string{
//   			jsii.String("timeout"),
//   		},
//   		UnhealthyThreshold: []*string{
//   			jsii.String("unhealthyThreshold"),
//   		},
//   	},
//   	Instances: []ResponseElementsItem{
//   		&ResponseElementsItem{
//   			InstanceId: []*string{
//   				jsii.String("instanceId"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// dNSName property.
	//
	// Specify an array of string values to match this event if the actual value of dNSName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DNsName *[]*string `field:"optional" json:"dNsName" yaml:"dNsName"`
	// healthCheck property.
	//
	// Specify an array of string values to match this event if the actual value of healthCheck is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HealthCheck *AWSAPICallViaCloudTrail_HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// instances property.
	//
	// Specify an array of string values to match this event if the actual value of instances is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Instances *[]*AWSAPICallViaCloudTrail_ResponseElementsItem `field:"optional" json:"instances" yaml:"instances"`
}

