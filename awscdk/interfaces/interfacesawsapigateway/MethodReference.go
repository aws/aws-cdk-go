package interfacesawsapigateway


// A reference to a Method resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   methodReference := &MethodReference{
//   	HttpMethod: jsii.String("httpMethod"),
//   	ResourceId: jsii.String("resourceId"),
//   	RestApiId: jsii.String("restApiId"),
//   }
//
type MethodReference struct {
	// The HttpMethod of the Method resource.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The ResourceId of the Method resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The RestApiId of the Method resource.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

