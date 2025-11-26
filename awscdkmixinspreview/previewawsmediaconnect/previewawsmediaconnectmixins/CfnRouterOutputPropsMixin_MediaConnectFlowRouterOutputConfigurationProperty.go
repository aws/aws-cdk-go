package previewawsmediaconnectmixins


// Configuration settings for connecting a router output to a MediaConnect flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var automatic interface{}
//
//   mediaConnectFlowRouterOutputConfigurationProperty := &MediaConnectFlowRouterOutputConfigurationProperty{
//   	DestinationTransitEncryption: &FlowTransitEncryptionProperty{
//   		EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		EncryptionKeyType: jsii.String("encryptionKeyType"),
//   	},
//   	FlowArn: jsii.String("flowArn"),
//   	FlowSourceArn: jsii.String("flowSourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-mediaconnectflowrouteroutputconfiguration.html
//
type CfnRouterOutputPropsMixin_MediaConnectFlowRouterOutputConfigurationProperty struct {
	// The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-mediaconnectflowrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-mediaconnectflowrouteroutputconfiguration-destinationtransitencryption
	//
	DestinationTransitEncryption interface{} `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
	// The ARN of the flow to connect to this router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-mediaconnectflowrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-mediaconnectflowrouteroutputconfiguration-flowarn
	//
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The ARN of the flow source to connect to this router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-mediaconnectflowrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-mediaconnectflowrouteroutputconfiguration-flowsourcearn
	//
	FlowSourceArn *string `field:"optional" json:"flowSourceArn" yaml:"flowSourceArn"`
}

