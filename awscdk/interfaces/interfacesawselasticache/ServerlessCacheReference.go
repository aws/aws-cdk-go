package interfacesawselasticache


// A reference to a ServerlessCache resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessCacheReference := &ServerlessCacheReference{
//   	ServerlessCacheArn: jsii.String("serverlessCacheArn"),
//   	ServerlessCacheName: jsii.String("serverlessCacheName"),
//   }
//
type ServerlessCacheReference struct {
	// The ARN of the ServerlessCache resource.
	ServerlessCacheArn *string `field:"required" json:"serverlessCacheArn" yaml:"serverlessCacheArn"`
	// The ServerlessCacheName of the ServerlessCache resource.
	ServerlessCacheName *string `field:"required" json:"serverlessCacheName" yaml:"serverlessCacheName"`
}

