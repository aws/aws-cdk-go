package awsrekognition


// The input parameters used to recognize faces in a streaming video analyzed by a Amazon Rekognition stream processor.
//
// `FaceSearchSettings` is a request parameter for [CreateStreamProcessor](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateStreamProcessor) . For more information, see [FaceSearchSettings](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_FaceSearchSettings) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   faceSearchSettingsProperty := &faceSearchSettingsProperty{
//   	collectionId: jsii.String("collectionId"),
//
//   	// the properties below are optional
//   	faceMatchThreshold: jsii.Number(123),
//   }
//
type CfnStreamProcessor_FaceSearchSettingsProperty struct {
	// The ID of a collection that contains faces that you want to search for.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// Minimum face match confidence score that must be met to return a result for a recognized face.
	//
	// The default is 80. 0 is the lowest confidence. 100 is the highest confidence. Values between 0 and 100 are accepted, and values lower than 80 are set to 80.
	FaceMatchThreshold *float64 `field:"optional" json:"faceMatchThreshold" yaml:"faceMatchThreshold"`
}

