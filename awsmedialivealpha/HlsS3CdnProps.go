package awsmedialivealpha


// Properties for HLS S3 CDN settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var s3CannedAcl S3CannedAcl
//
//   hlsS3CdnProps := &HlsS3CdnProps{
//   	CannedAcl: s3CannedAcl,
//   }
//
// Experimental.
type HlsS3CdnProps struct {
	// The S3 canned ACL to apply to each output.
	// Default: - no canned ACL.
	//
	// Experimental.
	CannedAcl S3CannedAcl `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
}

