package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configures a source monitoring metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringMetric := &MonitoringMetric{
//   	State: mediaconnect_alpha.State_ENABLED,
//   	Threshold: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type MonitoringMetric struct {
	// Whether the metric is enabled or disabled.
	//
	// The threshold is persisted by the
	// service either way, so you can toggle the metric on and off without losing
	// the threshold value.
	// Experimental.
	State State `field:"required" json:"state" yaml:"state"`
	// Threshold in seconds that triggers the metric's alert.
	//
	// Valid range: 10–60 seconds.
	// Experimental.
	Threshold awscdk.Duration `field:"required" json:"threshold" yaml:"threshold"`
}

