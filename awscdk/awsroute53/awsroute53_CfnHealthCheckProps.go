package awsroute53


// Properties for defining a `CfnHealthCheck`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHealthCheckProps := &cfnHealthCheckProps{
//   	healthCheckConfig: &healthCheckConfigProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		alarmIdentifier: &alarmIdentifierProperty{
//   			name: jsii.String("name"),
//   			region: jsii.String("region"),
//   		},
//   		childHealthChecks: []*string{
//   			jsii.String("childHealthChecks"),
//   		},
//   		enableSni: jsii.Boolean(false),
//   		failureThreshold: jsii.Number(123),
//   		fullyQualifiedDomainName: jsii.String("fullyQualifiedDomainName"),
//   		healthThreshold: jsii.Number(123),
//   		insufficientDataHealthStatus: jsii.String("insufficientDataHealthStatus"),
//   		inverted: jsii.Boolean(false),
//   		ipAddress: jsii.String("ipAddress"),
//   		measureLatency: jsii.Boolean(false),
//   		port: jsii.Number(123),
//   		regions: []*string{
//   			jsii.String("regions"),
//   		},
//   		requestInterval: jsii.Number(123),
//   		resourcePath: jsii.String("resourcePath"),
//   		routingControlArn: jsii.String("routingControlArn"),
//   		searchString: jsii.String("searchString"),
//   	},
//
//   	// the properties below are optional
//   	healthCheckTags: []interface{}{
//   		&healthCheckTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnHealthCheckProps struct {
	// A complex type that contains detailed information about one health check.
	//
	// For the values to enter for `HealthCheckConfig` , see [HealthCheckConfig](https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html)
	HealthCheckConfig interface{} `field:"required" json:"healthCheckConfig" yaml:"healthCheckConfig"`
	// The `HealthCheckTags` property describes key-value pairs that are associated with an `AWS::Route53::HealthCheck` resource.
	HealthCheckTags interface{} `field:"optional" json:"healthCheckTags" yaml:"healthCheckTags"`
}

