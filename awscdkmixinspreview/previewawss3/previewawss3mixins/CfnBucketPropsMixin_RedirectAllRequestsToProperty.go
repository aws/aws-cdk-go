package previewawss3mixins


// Specifies the redirect behavior of all requests to a website endpoint of an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redirectAllRequestsToProperty := &RedirectAllRequestsToProperty{
//   	HostName: jsii.String("hostName"),
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-redirectallrequeststo.html
//
type CfnBucketPropsMixin_RedirectAllRequestsToProperty struct {
	// Name of the host where requests are redirected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-redirectallrequeststo.html#cfn-s3-bucket-redirectallrequeststo-hostname
	//
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Protocol to use when redirecting requests.
	//
	// The default is the protocol that is used in the original request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-redirectallrequeststo.html#cfn-s3-bucket-redirectallrequeststo-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

