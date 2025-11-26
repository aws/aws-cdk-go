package previewawsredshiftmixins


// Describes a connection endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   endpointProperty := &EndpointProperty{
//   	Address: jsii.String("address"),
//   	Port: jsii.String("port"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html
//
type CfnClusterPropsMixin_EndpointProperty struct {
	// The DNS address of the cluster.
	//
	// This property is read only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The port that the database engine is listening on.
	//
	// This property is read only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-port
	//
	Port *string `field:"optional" json:"port" yaml:"port"`
}

