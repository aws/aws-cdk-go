package previewawsmediaconnectmixins


// Specifies the configuration settings for NDI outputs.
//
// Required when the flow includes NDI outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ndiConfigProperty := &NdiConfigProperty{
//   	MachineName: jsii.String("machineName"),
//   	NdiDiscoveryServers: []interface{}{
//   		&NdiDiscoveryServerConfigProperty{
//   			DiscoveryServerAddress: jsii.String("discoveryServerAddress"),
//   			DiscoveryServerPort: jsii.Number(123),
//   			VpcInterfaceAdapter: jsii.String("vpcInterfaceAdapter"),
//   		},
//   	},
//   	NdiState: jsii.String("ndiState"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndiconfig.html
//
type CfnFlowPropsMixin_NdiConfigProperty struct {
	// A prefix for the names of the NDI sources that the flow creates.
	//
	// If a custom name isn't specified, MediaConnect generates a unique 12-character ID as the prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndiconfig.html#cfn-mediaconnect-flow-ndiconfig-machinename
	//
	MachineName *string `field:"optional" json:"machineName" yaml:"machineName"`
	// A list of up to three NDI discovery server configurations.
	//
	// While not required by the API, this configuration is necessary for NDI functionality to work properly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndiconfig.html#cfn-mediaconnect-flow-ndiconfig-ndidiscoveryservers
	//
	NdiDiscoveryServers interface{} `field:"optional" json:"ndiDiscoveryServers" yaml:"ndiDiscoveryServers"`
	// A setting that controls whether NDI outputs can be used in the flow.
	//
	// Must be ENABLED to add NDI outputs. Default is DISABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndiconfig.html#cfn-mediaconnect-flow-ndiconfig-ndistate
	//
	NdiState *string `field:"optional" json:"ndiState" yaml:"ndiState"`
}

