package awss3


// Options for creating Virtual-Hosted style URL.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   bucket.urlForObject(jsii.String("objectname")) // Path-Style URL
//   bucket.virtualHostedUrlForObject(jsii.String("objectname")) // Virtual Hosted-Style URL
//   bucket.virtualHostedUrlForObject(jsii.String("objectname"), &virtualHostedStyleUrlOptions{
//   	regional: jsii.Boolean(false),
//   })
//
type VirtualHostedStyleUrlOptions struct {
	// Specifies the URL includes the region.
	Regional *bool `field:"optional" json:"regional" yaml:"regional"`
}

