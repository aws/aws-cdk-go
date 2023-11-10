package awscdkappconfigalpha


// The type of Monitor.
// Experimental.
type MonitorType string

const (
	// A Monitor from a CloudWatch alarm.
	// Experimental.
	MonitorType_CLOUDWATCH MonitorType = "CLOUDWATCH"
	// A Monitor from a CfnEnvironment.MonitorsProperty construct.
	// Experimental.
	MonitorType_CFN_MONITORS_PROPERTY MonitorType = "CFN_MONITORS_PROPERTY"
)

