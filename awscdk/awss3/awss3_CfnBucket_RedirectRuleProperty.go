package awss3


// Specifies how requests are redirected.
//
// In the event of an error, you can specify a different error code to return.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redirectRuleProperty := &redirectRuleProperty{
//   	hostName: jsii.String("hostName"),
//   	httpRedirectCode: jsii.String("httpRedirectCode"),
//   	protocol: jsii.String("protocol"),
//   	replaceKeyPrefixWith: jsii.String("replaceKeyPrefixWith"),
//   	replaceKeyWith: jsii.String("replaceKeyWith"),
//   }
//
type CfnBucket_RedirectRuleProperty struct {
	// The host name to use in the redirect request.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// The HTTP redirect code to use on the response.
	//
	// Not required if one of the siblings is present.
	HttpRedirectCode *string `field:"optional" json:"httpRedirectCode" yaml:"httpRedirectCode"`
	// Protocol to use when redirecting requests.
	//
	// The default is the protocol that is used in the original request.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The object key prefix to use in the redirect request.
	//
	// For example, to redirect requests for all pages with prefix `docs/` (objects in the `docs/` folder) to `documents/` , you can set a condition block with `KeyPrefixEquals` set to `docs/` and in the Redirect set `ReplaceKeyPrefixWith` to `/documents` . Not required if one of the siblings is present. Can be present only if `ReplaceKeyWith` is not provided.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	ReplaceKeyPrefixWith *string `field:"optional" json:"replaceKeyPrefixWith" yaml:"replaceKeyPrefixWith"`
	// The specific object key to use in the redirect request.
	//
	// For example, redirect request to `error.html` . Not required if one of the siblings is present. Can be present only if `ReplaceKeyPrefixWith` is not provided.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	ReplaceKeyWith *string `field:"optional" json:"replaceKeyWith" yaml:"replaceKeyWith"`
}

