package awsroute53


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigProperty := &healthCheckConfigProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	alarmIdentifier: &alarmIdentifierProperty{
//   		name: jsii.String("name"),
//   		region: jsii.String("region"),
//   	},
//   	childHealthChecks: []*string{
//   		jsii.String("childHealthChecks"),
//   	},
//   	enableSni: jsii.Boolean(false),
//   	failureThreshold: jsii.Number(123),
//   	fullyQualifiedDomainName: jsii.String("fullyQualifiedDomainName"),
//   	healthThreshold: jsii.Number(123),
//   	insufficientDataHealthStatus: jsii.String("insufficientDataHealthStatus"),
//   	inverted: jsii.Boolean(false),
//   	ipAddress: jsii.String("ipAddress"),
//   	measureLatency: jsii.Boolean(false),
//   	port: jsii.Number(123),
//   	regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	requestInterval: jsii.Number(123),
//   	resourcePath: jsii.String("resourcePath"),
//   	routingControlArn: jsii.String("routingControlArn"),
//   	searchString: jsii.String("searchString"),
//   }
//
type CfnHealthCheck_HealthCheckConfigProperty struct {
	// `CfnHealthCheck.HealthCheckConfigProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnHealthCheck.HealthCheckConfigProperty.AlarmIdentifier`.
	AlarmIdentifier interface{} `field:"optional" json:"alarmIdentifier" yaml:"alarmIdentifier"`
	// `CfnHealthCheck.HealthCheckConfigProperty.ChildHealthChecks`.
	ChildHealthChecks *[]*string `field:"optional" json:"childHealthChecks" yaml:"childHealthChecks"`
	// `CfnHealthCheck.HealthCheckConfigProperty.EnableSNI`.
	EnableSni interface{} `field:"optional" json:"enableSni" yaml:"enableSni"`
	// `CfnHealthCheck.HealthCheckConfigProperty.FailureThreshold`.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// `CfnHealthCheck.HealthCheckConfigProperty.FullyQualifiedDomainName`.
	FullyQualifiedDomainName *string `field:"optional" json:"fullyQualifiedDomainName" yaml:"fullyQualifiedDomainName"`
	// `CfnHealthCheck.HealthCheckConfigProperty.HealthThreshold`.
	HealthThreshold *float64 `field:"optional" json:"healthThreshold" yaml:"healthThreshold"`
	// `CfnHealthCheck.HealthCheckConfigProperty.InsufficientDataHealthStatus`.
	InsufficientDataHealthStatus *string `field:"optional" json:"insufficientDataHealthStatus" yaml:"insufficientDataHealthStatus"`
	// `CfnHealthCheck.HealthCheckConfigProperty.Inverted`.
	Inverted interface{} `field:"optional" json:"inverted" yaml:"inverted"`
	// `CfnHealthCheck.HealthCheckConfigProperty.IPAddress`.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// `CfnHealthCheck.HealthCheckConfigProperty.MeasureLatency`.
	MeasureLatency interface{} `field:"optional" json:"measureLatency" yaml:"measureLatency"`
	// `CfnHealthCheck.HealthCheckConfigProperty.Port`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// `CfnHealthCheck.HealthCheckConfigProperty.Regions`.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// `CfnHealthCheck.HealthCheckConfigProperty.RequestInterval`.
	RequestInterval *float64 `field:"optional" json:"requestInterval" yaml:"requestInterval"`
	// `CfnHealthCheck.HealthCheckConfigProperty.ResourcePath`.
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// `CfnHealthCheck.HealthCheckConfigProperty.RoutingControlArn`.
	RoutingControlArn *string `field:"optional" json:"routingControlArn" yaml:"routingControlArn"`
	// `CfnHealthCheck.HealthCheckConfigProperty.SearchString`.
	SearchString *string `field:"optional" json:"searchString" yaml:"searchString"`
}

