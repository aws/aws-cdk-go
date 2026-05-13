package awsdatazone


// S3 Properties Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3PropertiesInputProperty := &S3PropertiesInputProperty{
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	RegisterS3AccessGrantLocation: jsii.Boolean(false),
//   	S3AccessGrantLocationId: jsii.String("s3AccessGrantLocationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-s3propertiesinput.html
//
type CfnConnection_S3PropertiesInputProperty struct {
	// The Amazon S3 URI that's part of the Amazon S3 properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-s3propertiesinput.html#cfn-datazone-connection-s3propertiesinput-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Specifies whether to register the S3 Access Grant location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-s3propertiesinput.html#cfn-datazone-connection-s3propertiesinput-registers3accessgrantlocation
	//
	RegisterS3AccessGrantLocation interface{} `field:"optional" json:"registerS3AccessGrantLocation" yaml:"registerS3AccessGrantLocation"`
	// The Amazon S3 Access Grant location ID that's part of the Amazon S3 properties of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-s3propertiesinput.html#cfn-datazone-connection-s3propertiesinput-s3accessgrantlocationid
	//
	S3AccessGrantLocationId *string `field:"optional" json:"s3AccessGrantLocationId" yaml:"s3AccessGrantLocationId"`
}

