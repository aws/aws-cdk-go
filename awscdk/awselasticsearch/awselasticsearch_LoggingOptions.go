package awselasticsearch

import (
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
)

// Configures log settings for the domain.
//
// Example:
//   prodDomain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: es.elasticsearchVersion_V7_1(),
//   	capacity: &capacityConfig{
//   		masterNodes: jsii.Number(5),
//   		dataNodes: jsii.Number(20),
//   	},
//   	ebs: &ebsOptions{
//   		volumeSize: jsii.Number(20),
//   	},
//   	zoneAwareness: &zoneAwarenessConfig{
//   		availabilityZoneCount: jsii.Number(3),
//   	},
//   	logging: &loggingOptions{
//   		slowSearchLogEnabled: jsii.Boolean(true),
//   		appLogEnabled: jsii.Boolean(true),
//   		slowIndexLogEnabled: jsii.Boolean(true),
//   	},
//   })
//
// Deprecated: use opensearchservice module instead.
type LoggingOptions struct {
	// Specify if Elasticsearch application logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later.
	// Deprecated: use opensearchservice module instead.
	AppLogEnabled *bool `field:"optional" json:"appLogEnabled" yaml:"appLogEnabled"`
	// Log Elasticsearch application logs to this log group.
	// Deprecated: use opensearchservice module instead.
	AppLogGroup awslogs.ILogGroup `field:"optional" json:"appLogGroup" yaml:"appLogGroup"`
	// Specify if Elasticsearch audit logging should be set up.
	//
	// Requires Elasticsearch version 6.7 or later and fine grained access control to be enabled.
	// Deprecated: use opensearchservice module instead.
	AuditLogEnabled *bool `field:"optional" json:"auditLogEnabled" yaml:"auditLogEnabled"`
	// Log Elasticsearch audit logs to this log group.
	// Deprecated: use opensearchservice module instead.
	AuditLogGroup awslogs.ILogGroup `field:"optional" json:"auditLogGroup" yaml:"auditLogGroup"`
	// Specify if slow index logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later.
	// Deprecated: use opensearchservice module instead.
	SlowIndexLogEnabled *bool `field:"optional" json:"slowIndexLogEnabled" yaml:"slowIndexLogEnabled"`
	// Log slow indices to this log group.
	// Deprecated: use opensearchservice module instead.
	SlowIndexLogGroup awslogs.ILogGroup `field:"optional" json:"slowIndexLogGroup" yaml:"slowIndexLogGroup"`
	// Specify if slow search logging should be set up.
	//
	// Requires Elasticsearch version 5.1 or later.
	// Deprecated: use opensearchservice module instead.
	SlowSearchLogEnabled *bool `field:"optional" json:"slowSearchLogEnabled" yaml:"slowSearchLogEnabled"`
	// Log slow searches to this log group.
	// Deprecated: use opensearchservice module instead.
	SlowSearchLogGroup awslogs.ILogGroup `field:"optional" json:"slowSearchLogGroup" yaml:"slowSearchLogGroup"`
}

