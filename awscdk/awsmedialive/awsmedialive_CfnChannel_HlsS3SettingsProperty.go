package awsmedialive


// Sets up Amazon S3 as the destination for this HLS output.
//
// The parent of this entity is HlsCdnSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsS3SettingsProperty := &hlsS3SettingsProperty{
//   	cannedAcl: jsii.String("cannedAcl"),
//   }
//
type CfnChannel_HlsS3SettingsProperty struct {
	// Specify the canned ACL to apply to each S3 request.
	//
	// Defaults to none.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

