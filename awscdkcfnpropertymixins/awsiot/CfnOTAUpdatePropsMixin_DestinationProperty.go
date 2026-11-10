package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   destinationProperty := &DestinationProperty{
//   	S3Destination: &S3DestinationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-destination.html
//
type CfnOTAUpdatePropsMixin_DestinationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-destination.html#cfn-iot-otaupdate-destination-s3destination
	//
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

