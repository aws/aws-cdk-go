package previewawselasticloadbalancingevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	HealthCheck: &HealthCheck1{
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
//   	Listeners: []RequestParametersItem{
//   		&RequestParametersItem{
//   			InstancePort: []*string{
//   				jsii.String("instancePort"),
//   			},
//   			LoadBalancerPort: []*string{
//   				jsii.String("loadBalancerPort"),
//   			},
//   			Protocol: []*string{
//   				jsii.String("protocol"),
//   			},
//   		},
//   	},
//   	LoadBalancerName: []*string{
//   		jsii.String("loadBalancerName"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	TargetGroupArn: []*string{
//   		jsii.String("targetGroupArn"),
//   	},
//   	Targets: []RequestParametersItem1{
//   		&RequestParametersItem1{
//   			Id: []*string{
//   				jsii.String("id"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParameters struct {
	// healthCheck property.
	//
	// Specify an array of string values to match this event if the actual value of healthCheck is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HealthCheck *AWSAPICallViaCloudTrail_HealthCheck1 `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// instances property.
	//
	// Specify an array of string values to match this event if the actual value of instances is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Instances *[]*AWSAPICallViaCloudTrail_ResponseElementsItem `field:"optional" json:"instances" yaml:"instances"`
	// listeners property.
	//
	// Specify an array of string values to match this event if the actual value of listeners is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Listeners *[]*AWSAPICallViaCloudTrail_RequestParametersItem `field:"optional" json:"listeners" yaml:"listeners"`
	// loadBalancerName property.
	//
	// Specify an array of string values to match this event if the actual value of loadBalancerName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LoadBalancerName *[]*string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// subnets property.
	//
	// Specify an array of string values to match this event if the actual value of subnets is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// targetGroupArn property.
	//
	// Specify an array of string values to match this event if the actual value of targetGroupArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetGroupArn *[]*string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
	// targets property.
	//
	// Specify an array of string values to match this event if the actual value of targets is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Targets *[]*AWSAPICallViaCloudTrail_RequestParametersItem1 `field:"optional" json:"targets" yaml:"targets"`
}

