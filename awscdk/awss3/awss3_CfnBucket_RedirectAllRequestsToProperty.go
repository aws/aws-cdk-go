package awss3


// Specifies the redirect behavior of all requests to a website endpoint of an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redirectAllRequestsToProperty := &redirectAllRequestsToProperty{
//   	hostName: jsii.String("hostName"),
//
//   	// the properties below are optional
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnBucket_RedirectAllRequestsToProperty struct {
	// Name of the host where requests are redirected.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Protocol to use when redirecting requests.
	//
	// The default is the protocol that is used in the original request.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

