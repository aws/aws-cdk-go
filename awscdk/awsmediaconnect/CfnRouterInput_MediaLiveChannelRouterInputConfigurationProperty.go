package awsmediaconnect


// Configuration settings for connecting a router input to a MediaLive channel output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   mediaLiveChannelRouterInputConfigurationProperty := &MediaLiveChannelRouterInputConfigurationProperty{
//   	SourceTransitDecryption: &MediaLiveTransitEncryptionProperty{
//   		EncryptionKeyConfiguration: &MediaLiveTransitEncryptionKeyConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		EncryptionKeyType: jsii.String("encryptionKeyType"),
//   	},
//
//   	// the properties below are optional
//   	MediaLiveChannelArn: jsii.String("mediaLiveChannelArn"),
//   	MediaLiveChannelOutputName: jsii.String("mediaLiveChannelOutputName"),
//   	MediaLivePipelineId: jsii.String("mediaLivePipelineId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-medialivechannelrouterinputconfiguration.html
//
type CfnRouterInput_MediaLiveChannelRouterInputConfigurationProperty struct {
	// The encryption configuration that defines how content is encrypted during transit between MediaConnect Router and MediaLive.
	//
	// This configuration determines whether encryption keys are automatically managed by the service or manually managed through Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-medialivechannelrouterinputconfiguration.html#cfn-mediaconnect-routerinput-medialivechannelrouterinputconfiguration-sourcetransitdecryption
	//
	SourceTransitDecryption interface{} `field:"required" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
	// The ARN of the MediaLive channel to connect to this router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-medialivechannelrouterinputconfiguration.html#cfn-mediaconnect-routerinput-medialivechannelrouterinputconfiguration-medialivechannelarn
	//
	MediaLiveChannelArn *string `field:"optional" json:"mediaLiveChannelArn" yaml:"mediaLiveChannelArn"`
	// The name of the MediaLive channel output to connect to this router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-medialivechannelrouterinputconfiguration.html#cfn-mediaconnect-routerinput-medialivechannelrouterinputconfiguration-medialivechanneloutputname
	//
	MediaLiveChannelOutputName *string `field:"optional" json:"mediaLiveChannelOutputName" yaml:"mediaLiveChannelOutputName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-medialivechannelrouterinputconfiguration.html#cfn-mediaconnect-routerinput-medialivechannelrouterinputconfiguration-medialivepipelineid
	//
	MediaLivePipelineId *string `field:"optional" json:"mediaLivePipelineId" yaml:"mediaLivePipelineId"`
}

