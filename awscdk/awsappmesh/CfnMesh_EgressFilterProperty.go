package awsappmesh


// An object that represents the egress filter rules for a service mesh.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   egressFilterProperty := &EgressFilterProperty{
//   	Type: jsii.String("type"),
//   }
//
type CfnMesh_EgressFilterProperty struct {
	// The egress filter type.
	//
	// By default, the type is `DROP_ALL` , which allows egress only from virtual nodes to other defined resources in the service mesh (and any traffic to `*.amazonaws.com` for AWS API calls). You can set the egress filter type to `ALLOW_ALL` to allow egress to any endpoint inside or outside of the service mesh.
	//
	// > If you specify any backends on a virtual node when using `ALLOW_ALL` , you must specifiy all egress for that virtual node as backends. Otherwise, `ALLOW_ALL` will no longer work for that virtual node.
	Type *string `field:"required" json:"type" yaml:"type"`
}

