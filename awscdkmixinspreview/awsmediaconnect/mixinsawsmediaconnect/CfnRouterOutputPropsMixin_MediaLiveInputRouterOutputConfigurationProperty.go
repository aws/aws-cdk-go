package mixinsawsmediaconnect


// Configuration settings for connecting a router output to a MediaLive input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var automatic interface{}
//
//   mediaLiveInputRouterOutputConfigurationProperty := &MediaLiveInputRouterOutputConfigurationProperty{
//   	DestinationTransitEncryption: &MediaLiveTransitEncryptionProperty{
//   		EncryptionKeyConfiguration: &MediaLiveTransitEncryptionKeyConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		EncryptionKeyType: jsii.String("encryptionKeyType"),
//   	},
//   	MediaLiveInputArn: jsii.String("mediaLiveInputArn"),
//   	MediaLivePipelineId: jsii.String("mediaLivePipelineId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-medialiveinputrouteroutputconfiguration.html
//
type CfnRouterOutputPropsMixin_MediaLiveInputRouterOutputConfigurationProperty struct {
	// The encryption configuration that defines how content is encrypted during transit between MediaConnect Router and MediaLive.
	//
	// This configuration determines whether encryption keys are automatically managed by the service or manually managed through AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-medialiveinputrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-medialiveinputrouteroutputconfiguration-destinationtransitencryption
	//
	DestinationTransitEncryption interface{} `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
	// The ARN of the MediaLive input to connect to this router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-medialiveinputrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-medialiveinputrouteroutputconfiguration-medialiveinputarn
	//
	MediaLiveInputArn *string `field:"optional" json:"mediaLiveInputArn" yaml:"mediaLiveInputArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-medialiveinputrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-medialiveinputrouteroutputconfiguration-medialivepipelineid
	//
	MediaLivePipelineId *string `field:"optional" json:"mediaLivePipelineId" yaml:"mediaLivePipelineId"`
}

