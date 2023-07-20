package awsappmesh


// An object that represents the specification of a service mesh.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   meshSpecProperty := &MeshSpecProperty{
//   	EgressFilter: &EgressFilterProperty{
//   		Type: jsii.String("type"),
//   	},
//   	ServiceDiscovery: &MeshServiceDiscoveryProperty{
//   		IpPreference: jsii.String("ipPreference"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-mesh-meshspec.html
//
type CfnMesh_MeshSpecProperty struct {
	// The egress filter rules for the service mesh.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-mesh-meshspec.html#cfn-appmesh-mesh-meshspec-egressfilter
	//
	EgressFilter interface{} `field:"optional" json:"egressFilter" yaml:"egressFilter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-mesh-meshspec.html#cfn-appmesh-mesh-meshspec-servicediscovery
	//
	ServiceDiscovery interface{} `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

