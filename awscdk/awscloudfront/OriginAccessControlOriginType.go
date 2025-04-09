package awscloudfront


// Origin types supported by Origin Access Control.
type OriginAccessControlOriginType string

const (
	// Uses an Amazon S3 bucket origin.
	OriginAccessControlOriginType_S3 OriginAccessControlOriginType = "S3"
	// Uses a Lambda function URL origin.
	OriginAccessControlOriginType_LAMBDA OriginAccessControlOriginType = "LAMBDA"
	// Uses an AWS Elemental MediaStore origin.
	OriginAccessControlOriginType_MEDIASTORE OriginAccessControlOriginType = "MEDIASTORE"
	// Uses an AWS Elemental MediaPackage v2 origin.
	OriginAccessControlOriginType_MEDIAPACKAGEV2 OriginAccessControlOriginType = "MEDIAPACKAGEV2"
)

