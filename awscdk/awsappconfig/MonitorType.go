package awsappconfig


// The type of Monitor.
type MonitorType string

const (
	// A Monitor from a CloudWatch alarm.
	MonitorType_CLOUDWATCH MonitorType = "CLOUDWATCH"
	// A Monitor from a CfnEnvironment.MonitorsProperty construct.
	MonitorType_CFN_MONITORS_PROPERTY MonitorType = "CFN_MONITORS_PROPERTY"
)

