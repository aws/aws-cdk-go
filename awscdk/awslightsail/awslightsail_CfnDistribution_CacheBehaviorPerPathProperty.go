package awslightsail


// `CacheBehaviorPerPath` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the per-path cache behavior of an Amazon Lightsail content delivery network (CDN) distribution.
//
// Use a per-path cache behavior to override the default cache behavior of a distribution, or to add an exception to it. For example, if you set the `CacheBehavior` to `cache` , you can use a per-path cache behavior to specify a directory, file, or file type that your distribution will cache. If you don’t want your distribution to cache a specified directory, file, or file type, set the per-path cache behavior to `dont-cache` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheBehaviorPerPathProperty := &cacheBehaviorPerPathProperty{
//   	behavior: jsii.String("behavior"),
//   	path: jsii.String("path"),
//   }
//
type CfnDistribution_CacheBehaviorPerPathProperty struct {
	// The cache behavior for the specified path.
	//
	// You can specify one of the following per-path cache behaviors:
	//
	// - *`cache`* - This behavior caches the specified path.
	// - *`dont-cache`* - This behavior doesn't cache the specified path.
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// The path to a directory or file to cache, or not cache.
	//
	// Use an asterisk symbol to specify wildcard directories ( `path/to/assets/*` ), and file types ( `*.html` , `*jpg` , `*js` ). Directories and file paths are case-sensitive.
	//
	// Examples:
	//
	// - Specify the following to cache all files in the document root of an Apache web server running on a instance.
	//
	// `var/www/html/`
	// - Specify the following file to cache only the index page in the document root of an Apache web server.
	//
	// `var/www/html/index.html`
	// - Specify the following to cache only the .html files in the document root of an Apache web server.
	//
	// `var/www/html/*.html`
	// - Specify the following to cache only the .jpg, .png, and .gif files in the images sub-directory of the document root of an Apache web server.
	//
	// `var/www/html/images/*.jpg`
	//
	// `var/www/html/images/*.png`
	//
	// `var/www/html/images/*.gif`
	//
	// Specify the following to cache all files in the images subdirectory of the document root of an Apache web server.
	//
	// `var/www/html/images/`.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

