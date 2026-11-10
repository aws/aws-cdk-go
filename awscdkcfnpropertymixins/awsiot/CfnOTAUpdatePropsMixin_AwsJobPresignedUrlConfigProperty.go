package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   awsJobPresignedUrlConfigProperty := &AwsJobPresignedUrlConfigProperty{
//   	ExpiresInSec: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobpresignedurlconfig.html
//
type CfnOTAUpdatePropsMixin_AwsJobPresignedUrlConfigProperty struct {
	// How long (in seconds) pre-signed URLs are valid.
	//
	// Valid values are 60 - 3600, the default value is 1800 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobpresignedurlconfig.html#cfn-iot-otaupdate-awsjobpresignedurlconfig-expiresinsec
	//
	ExpiresInSec *float64 `field:"optional" json:"expiresInSec" yaml:"expiresInSec"`
}

