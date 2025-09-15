package awscdkamplifyalpha


// The type of cache configuration to use for an Amplify app.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
//   	CacheConfigType: amplify.CacheConfigType_AMPLIFY_MANAGED_NO_COOKIES,
//   })
//
// Experimental.
type CacheConfigType string

const (
	// AMPLIFY_MANAGED - Automatically applies an optimized cache configuration for your app based on its platform, routing rules, and rewrite rules.
	// Experimental.
	CacheConfigType_AMPLIFY_MANAGED CacheConfigType = "AMPLIFY_MANAGED"
	// AMPLIFY_MANAGED_NO_COOKIES - The same as AMPLIFY_MANAGED, except that it excludes all cookies from the cache key.
	// Experimental.
	CacheConfigType_AMPLIFY_MANAGED_NO_COOKIES CacheConfigType = "AMPLIFY_MANAGED_NO_COOKIES"
)

