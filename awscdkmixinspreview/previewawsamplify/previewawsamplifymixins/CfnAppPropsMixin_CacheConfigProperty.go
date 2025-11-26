package previewawsamplifymixins


// Describes the cache configuration for an Amplify app.
//
// For more information about how Amplify applies an optimal cache configuration for your app based on the type of content that is being served, see [Managing cache configuration](https://docs.aws.amazon.com/amplify/latest/userguide/managing-cache-configuration) in the *Amplify User guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cacheConfigProperty := &CacheConfigProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-cacheconfig.html
//
type CfnAppPropsMixin_CacheConfigProperty struct {
	// The type of cache configuration to use for an Amplify app.
	//
	// The `AMPLIFY_MANAGED` cache configuration automatically applies an optimized cache configuration for your app based on its platform, routing rules, and rewrite rules.
	//
	// The `AMPLIFY_MANAGED_NO_COOKIES` cache configuration type is the same as `AMPLIFY_MANAGED` , except that it excludes all cookies from the cache key. This is the default setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-cacheconfig.html#cfn-amplify-app-cacheconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

