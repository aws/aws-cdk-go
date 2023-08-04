package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Example:
//   var mesh mesh
//   // Cloud Map service discovery is currently required for host ejection by outlier detection
//   vpc := ec2.NewVpc(this, jsii.String("vpc"))
//   namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &PrivateDnsNamespaceProps{
//   	Vpc: Vpc,
//   	Name: jsii.String("domain.local"),
//   })
//   service := namespace.CreateService(jsii.String("Svc"))
//   node := mesh.addVirtualNode(jsii.String("virtual-node"), &VirtualNodeBaseProps{
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			OutlierDetection: &OutlierDetection{
//   				BaseEjectionDuration: awscdk.Duration_Seconds(jsii.Number(10)),
//   				Interval: awscdk.Duration_*Seconds(jsii.Number(30)),
//   				MaxEjectionPercent: jsii.Number(50),
//   				MaxServerErrors: jsii.Number(5),
//   			},
//   		}),
//   	},
//   })
//
type PrivateDnsNamespaceProps struct {
	// A name for the Namespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the Namespace.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon VPC that you want to associate the namespace with.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

