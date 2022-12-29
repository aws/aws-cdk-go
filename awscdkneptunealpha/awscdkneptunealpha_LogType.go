// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkneptunealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Neptune log types that can be exported to CloudWatch logs.
//
// Example:
//   // Cluster parameter group with the neptune_enable_audit_log param set to 1
//   clusterParameterGroup := neptune.NewClusterParameterGroup(this, jsii.String("ClusterParams"), &clusterParameterGroupProps{
//   	description: jsii.String("Cluster parameter group"),
//   	parameters: map[string]*string{
//   		"neptune_enable_audit_log": jsii.String("1"),
//   	},
//   })
//
//   cluster := neptune.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	vpc: vpc,
//   	instanceType: neptune.instanceType_R5_LARGE(),
//   	// Audit logs are enabled via the clusterParameterGroup
//   	clusterParameterGroup: clusterParameterGroup,
//   	// Optionally configuring audit logs to be exported to CloudWatch Logs
//   	cloudwatchLogsExports: []logType{
//   		neptune.*logType_AUDIT(),
//   	},
//   	// Optionally set a retention period on exported CloudWatch Logs
//   	cloudwatchLogsRetention: logs.retentionDays_ONE_MONTH,
//   })
//
// See: https://docs.aws.amazon.com/neptune/latest/userguide/cloudwatch-logs.html
//
// Experimental.
type LogType interface {
	// the log type.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for LogType
type jsiiProxy_LogType struct {
	_ byte // padding
}

func (j *jsiiProxy_LogType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Constructor for specifying a custom log type.
// Experimental.
func NewLogType(value *string) LogType {
	_init_.Initialize()

	if err := validateNewLogTypeParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogType{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.LogType",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Constructor for specifying a custom log type.
// Experimental.
func NewLogType_Override(l LogType, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.LogType",
		[]interface{}{value},
		l,
	)
}

func LogType_AUDIT() LogType {
	_init_.Initialize()
	var returns LogType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.LogType",
		"AUDIT",
		&returns,
	)
	return returns
}

