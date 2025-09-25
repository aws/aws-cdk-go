package awsapigateway


// A reference to a RestApi resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restApiReference := &RestApiReference{
//   	RestApiId: jsii.String("restApiId"),
//   }
//
type RestApiReference struct {
	// The RestApiId of the RestApi resource.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

