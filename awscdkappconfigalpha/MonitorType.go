package awscdkappconfigalpha


// The type of Monitor.
// Deprecated.
type MonitorType string

const (
	// A Monitor from a CloudWatch alarm.
	// Deprecated.
	MonitorType_CLOUDWATCH MonitorType = "CLOUDWATCH"
	// A Monitor from a CfnEnvironment.MonitorsProperty construct.
	// Deprecated.
	MonitorType_CFN_MONITORS_PROPERTY MonitorType = "CFN_MONITORS_PROPERTY"
)

