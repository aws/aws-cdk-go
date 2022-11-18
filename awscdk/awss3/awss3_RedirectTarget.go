package awss3


// Specifies a redirect behavior of all requests to a website endpoint of a bucket.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &bucketProps{
//   	websiteRedirect: &redirectTarget{
//   		hostName: jsii.String("www.example.com"),
//   	},
//   })
//
type RedirectTarget struct {
	// Name of the host where requests are redirected.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Protocol to use when redirecting requests.
	Protocol RedirectProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

