package awsmediaconnect


// Configuration settings for connecting a router input to a flow output.
//
// To connect the router input to the flow output, provide both flowArn and flowOutputArn. To disconnect it, call the update operation and omit both fields while keeping this configuration object. You can't change sourceTransitDecryption while the router input is connected to a flow output. To change it, first disconnect the router input, then update the decryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var automatic interface{}
//
//   mediaConnectFlowRouterInputConfigurationProperty := &MediaConnectFlowRouterInputConfigurationProperty{
//   	FlowArn: jsii.String("flowArn"),
//   	FlowOutputArn: jsii.String("flowOutputArn"),
//   	SourceTransitDecryption: &FlowTransitEncryptionProperty{
//   		EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		EncryptionKeyType: jsii.String("encryptionKeyType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mediaconnectflowrouterinputconfiguration.html
//
type CfnRouterInputPropsMixin_MediaConnectFlowRouterInputConfigurationProperty struct {
	// The ARN of the flow to connect to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mediaconnectflowrouterinputconfiguration.html#cfn-mediaconnect-routerinput-mediaconnectflowrouterinputconfiguration-flowarn
	//
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The ARN of the flow output to connect to this router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mediaconnectflowrouterinputconfiguration.html#cfn-mediaconnect-routerinput-mediaconnectflowrouterinputconfiguration-flowoutputarn
	//
	FlowOutputArn *string `field:"optional" json:"flowOutputArn" yaml:"flowOutputArn"`
	// The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-mediaconnectflowrouterinputconfiguration.html#cfn-mediaconnect-routerinput-mediaconnectflowrouterinputconfiguration-sourcetransitdecryption
	//
	SourceTransitDecryption interface{} `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
}

