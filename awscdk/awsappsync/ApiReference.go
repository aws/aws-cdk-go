package awsappsync


// A reference to a Api resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiReference := &ApiReference{
//   	ApiArn: jsii.String("apiArn"),
//   }
//
type ApiReference struct {
	// The ApiArn of the Api resource.
	ApiArn *string `field:"required" json:"apiArn" yaml:"apiArn"`
}

