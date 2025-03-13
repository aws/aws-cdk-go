package awsroute53


// Properties for defining a `CfnHealthCheck`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHealthCheckProps := &CfnHealthCheckProps{
//   	HealthCheckConfig: &HealthCheckConfigProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		AlarmIdentifier: &AlarmIdentifierProperty{
//   			Name: jsii.String("name"),
//   			Region: jsii.String("region"),
//   		},
//   		ChildHealthChecks: []*string{
//   			jsii.String("childHealthChecks"),
//   		},
//   		EnableSni: jsii.Boolean(false),
//   		FailureThreshold: jsii.Number(123),
//   		FullyQualifiedDomainName: jsii.String("fullyQualifiedDomainName"),
//   		HealthThreshold: jsii.Number(123),
//   		InsufficientDataHealthStatus: jsii.String("insufficientDataHealthStatus"),
//   		Inverted: jsii.Boolean(false),
//   		IpAddress: jsii.String("ipAddress"),
//   		MeasureLatency: jsii.Boolean(false),
//   		Port: jsii.Number(123),
//   		Regions: []*string{
//   			jsii.String("regions"),
//   		},
//   		RequestInterval: jsii.Number(123),
//   		ResourcePath: jsii.String("resourcePath"),
//   		RoutingControlArn: jsii.String("routingControlArn"),
//   		SearchString: jsii.String("searchString"),
//   	},
//
//   	// the properties below are optional
//   	HealthCheckTags: []interface{}{
//   		&HealthCheckTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
//
type CfnHealthCheckProps struct {
	// A complex type that contains detailed information about one health check.
	//
	// For the values to enter for `HealthCheckConfig` , see [HealthCheckConfig](https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig
	//
	HealthCheckConfig interface{} `field:"required" json:"healthCheckConfig" yaml:"healthCheckConfig"`
	// The `HealthCheckTags` property describes key-value pairs that are associated with an `AWS::Route53::HealthCheck` resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags
	//
	HealthCheckTags interface{} `field:"optional" json:"healthCheckTags" yaml:"healthCheckTags"`
}

