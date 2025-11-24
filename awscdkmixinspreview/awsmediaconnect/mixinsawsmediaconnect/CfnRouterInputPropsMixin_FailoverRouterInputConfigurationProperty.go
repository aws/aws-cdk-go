package mixinsawsmediaconnect


// Configuration settings for a failover router input that allows switching between two input sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   failoverRouterInputConfigurationProperty := &FailoverRouterInputConfigurationProperty{
//   	NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   	PrimarySourceIndex: jsii.Number(123),
//   	ProtocolConfigurations: []interface{}{
//   		&FailoverRouterInputProtocolConfigurationProperty{
//   			Rist: &RistRouterInputConfigurationProperty{
//   				Port: jsii.Number(123),
//   				RecoveryLatencyMilliseconds: jsii.Number(123),
//   			},
//   			Rtp: &RtpRouterInputConfigurationProperty{
//   				ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   				Port: jsii.Number(123),
//   			},
//   			SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   				DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				SourceAddress: jsii.String("sourceAddress"),
//   				SourcePort: jsii.Number(123),
//   				StreamId: jsii.String("streamId"),
//   			},
//   			SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   				DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SourcePriorityMode: jsii.String("sourcePriorityMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html
//
type CfnRouterInputPropsMixin_FailoverRouterInputConfigurationProperty struct {
	// The ARN of the network interface to use for this failover router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputconfiguration-networkinterfacearn
	//
	NetworkInterfaceArn *string `field:"optional" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// The index (0 or 1) that specifies which source in the protocol configurations list is currently active.
	//
	// Used to control which of the two failover sources is currently selected. This field is ignored when sourcePriorityMode is set to NO_PRIORITY
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputconfiguration-primarysourceindex
	//
	PrimarySourceIndex *float64 `field:"optional" json:"primarySourceIndex" yaml:"primarySourceIndex"`
	// A list of exactly two protocol configurations for the failover input sources.
	//
	// Both must use the same protocol type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputconfiguration-protocolconfigurations
	//
	ProtocolConfigurations interface{} `field:"optional" json:"protocolConfigurations" yaml:"protocolConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputconfiguration-sourceprioritymode
	//
	SourcePriorityMode *string `field:"optional" json:"sourcePriorityMode" yaml:"sourcePriorityMode"`
}

