package awsevs


// > Amazon EVS is in public preview release and is subject to change.
//
// The connectivity configuration for the environment. Amazon EVS requires that you specify two route server peer IDs. During environment creation, the route server endpoints peer with the NSX uplink VLAN for connectivity to the NSX overlay network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectivityInfoProperty := &ConnectivityInfoProperty{
//   	PrivateRouteServerPeerings: []*string{
//   		jsii.String("privateRouteServerPeerings"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-connectivityinfo.html
//
type CfnEnvironment_ConnectivityInfoProperty struct {
	// The unique IDs for private route server peers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-connectivityinfo.html#cfn-evs-environment-connectivityinfo-privaterouteserverpeerings
	//
	PrivateRouteServerPeerings *[]*string `field:"required" json:"privateRouteServerPeerings" yaml:"privateRouteServerPeerings"`
}

