package awsapigateway


// A reference to a Resource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceReference := &ResourceReference{
//   	ResourceId: jsii.String("resourceId"),
//   	RestApiId: jsii.String("restApiId"),
//   }
//
type ResourceReference struct {
	// The ResourceId of the Resource resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The RestApiId of the Resource resource.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

