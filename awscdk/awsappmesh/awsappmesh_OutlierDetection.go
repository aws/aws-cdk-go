package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Represents the outlier detection for a listener.
//
// Example:
//   var mesh mesh
//   // Cloud Map service discovery is currently required for host ejection by outlier detection
//   vpc := ec2.NewVpc(this, jsii.String("vpc"))
//   namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &privateDnsNamespaceProps{
//   	vpc: vpc,
//   	name: jsii.String("domain.local"),
//   })
//   service := namespace.createService(jsii.String("Svc"))
//   node := mesh.addVirtualNode(jsii.String("virtual-node"), &virtualNodeBaseProps{
//   	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			outlierDetection: &outlierDetection{
//   				baseEjectionDuration: cdk.duration.seconds(jsii.Number(10)),
//   				interval: cdk.*duration.seconds(jsii.Number(30)),
//   				maxEjectionPercent: jsii.Number(50),
//   				maxServerErrors: jsii.Number(5),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type OutlierDetection struct {
	// The base amount of time for which a host is ejected.
	// Experimental.
	BaseEjectionDuration awscdk.Duration `field:"required" json:"baseEjectionDuration" yaml:"baseEjectionDuration"`
	// The time interval between ejection sweep analysis.
	// Experimental.
	Interval awscdk.Duration `field:"required" json:"interval" yaml:"interval"`
	// Maximum percentage of hosts in load balancing pool for upstream service that can be ejected.
	//
	// Will eject at
	// least one host regardless of the value.
	// Experimental.
	MaxEjectionPercent *float64 `field:"required" json:"maxEjectionPercent" yaml:"maxEjectionPercent"`
	// Number of consecutive 5xx errors required for ejection.
	// Experimental.
	MaxServerErrors *float64 `field:"required" json:"maxServerErrors" yaml:"maxServerErrors"`
}

