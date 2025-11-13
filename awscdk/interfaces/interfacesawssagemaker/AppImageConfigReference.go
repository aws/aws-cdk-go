package interfacesawssagemaker


// A reference to a AppImageConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appImageConfigReference := &AppImageConfigReference{
//   	AppImageConfigArn: jsii.String("appImageConfigArn"),
//   	AppImageConfigName: jsii.String("appImageConfigName"),
//   }
//
type AppImageConfigReference struct {
	// The ARN of the AppImageConfig resource.
	AppImageConfigArn *string `field:"required" json:"appImageConfigArn" yaml:"appImageConfigArn"`
	// The AppImageConfigName of the AppImageConfig resource.
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
}

