package awsworkspaces


// Describes the access endpoint configuration for a WorkSpace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   accessEndpointConfigProperty := &AccessEndpointConfigProperty{
//   	AccessEndpoints: []interface{}{
//   		&AccessEndpointProperty{
//   			AccessEndpointType: jsii.String("accessEndpointType"),
//   			VpcEndpointId: jsii.String("vpcEndpointId"),
//   		},
//   	},
//   	InternetFallbackProtocols: []*string{
//   		jsii.String("internetFallbackProtocols"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-accessendpointconfig.html
//
type CfnDirectoryPropsMixin_AccessEndpointConfigProperty struct {
	// Indicates a list of access endpoints associated with this directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-accessendpointconfig.html#cfn-workspaces-directory-accessendpointconfig-accessendpoints
	//
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// Indicates a list of protocols that fallback to using the public Internet when streaming over a VPC endpoint is not available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-accessendpointconfig.html#cfn-workspaces-directory-accessendpointconfig-internetfallbackprotocols
	//
	InternetFallbackProtocols *[]*string `field:"optional" json:"internetFallbackProtocols" yaml:"internetFallbackProtocols"`
}

