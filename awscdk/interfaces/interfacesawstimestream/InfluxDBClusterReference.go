package interfacesawstimestream


// A reference to a InfluxDBCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   influxDBClusterReference := &InfluxDBClusterReference{
//   	InfluxDbClusterArn: jsii.String("influxDbClusterArn"),
//   	InfluxDbClusterId: jsii.String("influxDbClusterId"),
//   }
//
type InfluxDBClusterReference struct {
	// The ARN of the InfluxDBCluster resource.
	InfluxDbClusterArn *string `field:"required" json:"influxDbClusterArn" yaml:"influxDbClusterArn"`
	// The Id of the InfluxDBCluster resource.
	InfluxDbClusterId *string `field:"required" json:"influxDbClusterId" yaml:"influxDbClusterId"`
}

