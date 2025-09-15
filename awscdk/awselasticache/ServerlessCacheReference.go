package awselasticache


// A reference to a ServerlessCache resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessCacheReference := &ServerlessCacheReference{
//   	ServerlessCacheName: jsii.String("serverlessCacheName"),
//   }
//
type ServerlessCacheReference struct {
	// The ServerlessCacheName of the ServerlessCache resource.
	ServerlessCacheName *string `field:"required" json:"serverlessCacheName" yaml:"serverlessCacheName"`
}

