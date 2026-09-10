package interfacesawschime


// A reference to a MediaPipelineKinesisVideoStreamPool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaPipelineKinesisVideoStreamPoolReference := &MediaPipelineKinesisVideoStreamPoolReference{
//   	MediaPipelineKinesisVideoStreamPoolArn: jsii.String("mediaPipelineKinesisVideoStreamPoolArn"),
//   }
//
type MediaPipelineKinesisVideoStreamPoolReference struct {
	// The Arn of the MediaPipelineKinesisVideoStreamPool resource.
	MediaPipelineKinesisVideoStreamPoolArn *string `field:"required" json:"mediaPipelineKinesisVideoStreamPoolArn" yaml:"mediaPipelineKinesisVideoStreamPoolArn"`
}

