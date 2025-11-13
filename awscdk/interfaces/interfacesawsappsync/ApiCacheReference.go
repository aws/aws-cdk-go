package interfacesawsappsync


// A reference to a ApiCache resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiCacheReference := &ApiCacheReference{
//   	ApiCacheId: jsii.String("apiCacheId"),
//   }
//
type ApiCacheReference struct {
	// The Id of the ApiCache resource.
	ApiCacheId *string `field:"required" json:"apiCacheId" yaml:"apiCacheId"`
}

