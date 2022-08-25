package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type PrivateDnsNamespaceProps struct {
	// A name for the Namespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the Namespace.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon VPC that you want to associate the namespace with.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

