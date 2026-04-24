package awsecs


// Configuration for Service Connect access logs.
//
// Service Connect access logs provide detailed telemetry about individual requests processed by the Service Connect proxy,
// including HTTP methods, paths, response codes, and timing information.
//
// Example:
//   var cluster Cluster
//   var taskDefinition TaskDefinition
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	ServiceConnectConfiguration: &ServiceConnectProps{
//   		Services: []ServiceConnectService{
//   			&ServiceConnectService{
//   				PortMappingName: jsii.String("api"),
//   			},
//   		},
//   		AccessLogConfiguration: &ServiceConnectAccessLogConfiguration{
//   			Format: ecs.ServiceConnectAccessLogFormat_JSON,
//   			IncludeQueryParameters: jsii.Boolean(true),
//   		},
//   		// When configuring access log,
//   		// you also need to configure the log driver accordingly.
//   		LogDriver: ecs.LogDrivers_AwsLogs(&AwsLogDriverProps{
//   			StreamPrefix: jsii.String("prefix"),
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect-envoy-access-logs.html
//
type ServiceConnectAccessLogConfiguration struct {
	// The format for Service Connect access log output.
	//
	// - TEXT: Human-readable text format
	// - JSON: Structured JSON format for log analysis tools.
	Format ServiceConnectAccessLogFormat `field:"required" json:"format" yaml:"format"`
	// Whether to include query parameters in Service Connect access logs.
	//
	// When enabled, query parameters from HTTP requests are included in the access logs.
	// Consider security and privacy implications as query parameters may contain sensitive information such as request IDs and tokens.
	// Default: undefined - AWS ECS default is false, which means that query parameters are not included in access logs.
	//
	IncludeQueryParameters *bool `field:"optional" json:"includeQueryParameters" yaml:"includeQueryParameters"`
}

