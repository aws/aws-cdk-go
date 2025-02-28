package awss3


// Options for creating Virtual-Hosted style URL.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   bucket.UrlForObject(jsii.String("objectname")) // Path-Style URL
//   bucket.VirtualHostedUrlForObject(jsii.String("objectname")) // Virtual Hosted-Style URL
//   bucket.VirtualHostedUrlForObject(jsii.String("objectname"), &VirtualHostedStyleUrlOptions{
//   	Regional: jsii.Boolean(false),
//   })
//
type VirtualHostedStyleUrlOptions struct {
	// Specifies the URL includes the region.
	// Default: - true.
	//
	Regional *bool `field:"optional" json:"regional" yaml:"regional"`
}

