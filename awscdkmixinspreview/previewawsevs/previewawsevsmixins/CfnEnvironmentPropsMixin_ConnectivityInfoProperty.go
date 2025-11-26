package previewawsevsmixins


// The connectivity configuration for the environment.
//
// Amazon EVS requires that you specify two route server peer IDs. During environment creation, the route server endpoints peer with the NSX uplink VLAN for connectivity to the NSX overlay network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   connectivityInfoProperty := &ConnectivityInfoProperty{
//   	PrivateRouteServerPeerings: []*string{
//   		jsii.String("privateRouteServerPeerings"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-connectivityinfo.html
//
type CfnEnvironmentPropsMixin_ConnectivityInfoProperty struct {
	// The unique IDs for private route server peers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-connectivityinfo.html#cfn-evs-environment-connectivityinfo-privaterouteserverpeerings
	//
	PrivateRouteServerPeerings *[]*string `field:"optional" json:"privateRouteServerPeerings" yaml:"privateRouteServerPeerings"`
}

