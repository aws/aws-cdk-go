package awsmediaconnect


// Configuration settings for a failover router input that allows switching between two input sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverRouterInputConfigurationProperty := &FailoverRouterInputConfigurationProperty{
//   	NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   	ProtocolConfigurations: []interface{}{
//   		&FailoverRouterInputProtocolConfigurationProperty{
//   			Rist: &RistRouterInputConfigurationProperty{
//   				Port: jsii.Number(123),
//   				RecoveryLatencyMilliseconds: jsii.Number(123),
//   			},
//   			Rtp: &RtpRouterInputConfigurationProperty{
//   				Port: jsii.Number(123),
//
//   				// the properties below are optional
//   				ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   			},
//   			SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				SourceAddress: jsii.String("sourceAddress"),
//   				SourcePort: jsii.Number(123),
//
//   				// the properties below are optional
//   				DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   				StreamId: jsii.String("streamId"),
//   			},
//   			SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				Port: jsii.Number(123),
//
//   				// the properties below are optional
//   				DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   					EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						SecretArn: jsii.String("secretArn"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SourcePriorityMode: jsii.String("sourcePriorityMode"),
//
//   	// the properties below are optional
//   	PrimarySourceIndex: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html
//
type CfnRouterInput_FailoverRouterInputConfigurationProperty struct {
	// The ARN of the network interface to use for this failover router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputconfiguration-networkinterfacearn
	//
	NetworkInterfaceArn *string `field:"required" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// A list of exactly two protocol configurations for the failover input sources.
	//
	// Both must use the same protocol type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputconfiguration-protocolconfigurations
	//
	ProtocolConfigurations interface{} `field:"required" json:"protocolConfigurations" yaml:"protocolConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputconfiguration-sourceprioritymode
	//
	SourcePriorityMode *string `field:"required" json:"sourcePriorityMode" yaml:"sourcePriorityMode"`
	// The index (0 or 1) that specifies which source in the protocol configurations list is currently active.
	//
	// Used to control which of the two failover sources is currently selected. This field is ignored when sourcePriorityMode is set to NO_PRIORITY
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.html#cfn-mediaconnect-routerinput-failoverrouterinputconfiguration-primarysourceindex
	//
	PrimarySourceIndex *float64 `field:"optional" json:"primarySourceIndex" yaml:"primarySourceIndex"`
}

