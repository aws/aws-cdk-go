package awswafv2


// The filter to use to identify the subset of cookies to inspect in a web request.
//
// You must specify exactly one setting: either `All` , `IncludedCookies` , or `ExcludedCookies` .
//
// Example JSON: `"MatchPattern": { "IncludedCookies": [ "session-id-time", "session-id" ] }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   cookieMatchPatternProperty := &CookieMatchPatternProperty{
//   	All: all,
//   	ExcludedCookies: []*string{
//   		jsii.String("excludedCookies"),
//   	},
//   	IncludedCookies: []*string{
//   		jsii.String("includedCookies"),
//   	},
//   }
//
type CfnWebACL_CookieMatchPatternProperty struct {
	// Inspect all cookies.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Inspect only the cookies whose keys don't match any of the strings specified here.
	ExcludedCookies *[]*string `field:"optional" json:"excludedCookies" yaml:"excludedCookies"`
	// Inspect only the cookies that have a key that matches one of the strings specified here.
	IncludedCookies *[]*string `field:"optional" json:"includedCookies" yaml:"includedCookies"`
}

