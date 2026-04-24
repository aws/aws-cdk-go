package awsecs


// The format of Service Connect access logs.
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
type ServiceConnectAccessLogFormat string

const (
	// Human-readable text format for access logs.
	ServiceConnectAccessLogFormat_TEXT ServiceConnectAccessLogFormat = "TEXT"
	// Structured JSON format for access logs.
	//
	// This format is well-suited for integration with log analysis tools.
	ServiceConnectAccessLogFormat_JSON ServiceConnectAccessLogFormat = "JSON"
)

