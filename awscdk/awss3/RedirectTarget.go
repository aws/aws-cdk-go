package awss3


// Specifies a redirect behavior of all requests to a website endpoint of a bucket.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &BucketProps{
//   	WebsiteRedirect: &RedirectTarget{
//   		HostName: jsii.String("www.example.com"),
//   	},
//   })
//
type RedirectTarget struct {
	// Name of the host where requests are redirected.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Protocol to use when redirecting requests.
	// Default: - The protocol used in the original request.
	//
	Protocol RedirectProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

