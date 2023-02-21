// Version 2 of the AWS Cloud Development Kit library
package awscdk


// An enum representing the various ARN formats that different services use.
type ArnFormat string

const (
	// This represents a format where there is no 'resourceName' part.
	//
	// This format is used for S3 resources,
	// like 'arn:aws:s3:::bucket'.
	// Everything after the last colon is considered the 'resource',
	// even if it contains slashes,
	// like in 'arn:aws:s3:::bucket/object.zip'.
	ArnFormat_NO_RESOURCE_NAME ArnFormat = "NO_RESOURCE_NAME"
	// This represents a format where the 'resource' and 'resourceName' parts are separated with a colon.
	//
	// Like in: 'arn:aws:service:region:account:resource:resourceName'.
	// Everything after the last colon is considered the 'resourceName',
	// even if it contains slashes,
	// like in 'arn:aws:apigateway:region:account:resource:/test/mydemoresource/*'.
	ArnFormat_COLON_RESOURCE_NAME ArnFormat = "COLON_RESOURCE_NAME"
	// This represents a format where the 'resource' and 'resourceName' parts are separated with a slash.
	//
	// Like in: 'arn:aws:service:region:account:resource/resourceName'.
	// Everything after the separating slash is considered the 'resourceName',
	// even if it contains colons,
	// like in 'arn:aws:cognito-sync:region:account:identitypool/us-east-1:1a1a1a1a-ffff-1111-9999-12345678:bla'.
	ArnFormat_SLASH_RESOURCE_NAME ArnFormat = "SLASH_RESOURCE_NAME"
	// This represents a format where the 'resource' and 'resourceName' parts are seperated with a slash, but there is also an additional slash after the colon separating 'account' from 'resource'.
	//
	// Like in: 'arn:aws:service:region:account:/resource/resourceName'.
	// Note that the leading slash is _not_ included in the parsed 'resource' part.
	ArnFormat_SLASH_RESOURCE_SLASH_RESOURCE_NAME ArnFormat = "SLASH_RESOURCE_SLASH_RESOURCE_NAME"
)

