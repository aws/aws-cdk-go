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
// Experimental.
type RedirectTarget struct {
	// Name of the host where requests are redirected.
	// Experimental.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Protocol to use when redirecting requests.
	// Experimental.
	Protocol RedirectProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

