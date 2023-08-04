package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Configures log settings for the domain.
//
// Example:
//   prodDomain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	Capacity: &CapacityConfig{
//   		MasterNodes: jsii.Number(5),
//   		DataNodes: jsii.Number(20),
//   	},
//   	Ebs: &EbsOptions{
//   		VolumeSize: jsii.Number(20),
//   	},
//   	ZoneAwareness: &ZoneAwarenessConfig{
//   		AvailabilityZoneCount: jsii.Number(3),
//   	},
//   	Logging: &LoggingOptions{
//   		SlowSearchLogEnabled: jsii.Boolean(true),
//   		AppLogEnabled: jsii.Boolean(true),
//   		SlowIndexLogEnabled: jsii.Boolean(true),
//   	},
//   })
//
type LoggingOptions struct {
	// Specify if Amazon OpenSearch Service application logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later or OpenSearch version 1.0 or later.
	// Default: - false.
	//
	AppLogEnabled *bool `field:"optional" json:"appLogEnabled" yaml:"appLogEnabled"`
	// Log Amazon OpenSearch Service application logs to this log group.
	// Default: - a new log group is created if app logging is enabled.
	//
	AppLogGroup awslogs.ILogGroup `field:"optional" json:"appLogGroup" yaml:"appLogGroup"`
	// Specify if Amazon OpenSearch Service audit logging should be set up.
	//
	// Requires Elasticsearch version 6.7 or later or OpenSearch version 1.0 or later and fine grained access control to be enabled.
	// Default: - false.
	//
	AuditLogEnabled *bool `field:"optional" json:"auditLogEnabled" yaml:"auditLogEnabled"`
	// Log Amazon OpenSearch Service audit logs to this log group.
	// Default: - a new log group is created if audit logging is enabled.
	//
	AuditLogGroup awslogs.ILogGroup `field:"optional" json:"auditLogGroup" yaml:"auditLogGroup"`
	// Specify if slow index logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later or OpenSearch version 1.0 or later.
	// Default: - false.
	//
	SlowIndexLogEnabled *bool `field:"optional" json:"slowIndexLogEnabled" yaml:"slowIndexLogEnabled"`
	// Log slow indices to this log group.
	// Default: - a new log group is created if slow index logging is enabled.
	//
	SlowIndexLogGroup awslogs.ILogGroup `field:"optional" json:"slowIndexLogGroup" yaml:"slowIndexLogGroup"`
	// Specify if slow search logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later or OpenSearch version 1.0 or later.
	// Default: - false.
	//
	SlowSearchLogEnabled *bool `field:"optional" json:"slowSearchLogEnabled" yaml:"slowSearchLogEnabled"`
	// Log slow searches to this log group.
	// Default: - a new log group is created if slow search logging is enabled.
	//
	SlowSearchLogGroup awslogs.ILogGroup `field:"optional" json:"slowSearchLogGroup" yaml:"slowSearchLogGroup"`
}

