package interfacesawsroute53


// A reference to a QueryLoggingConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryLoggingConfigReference := &QueryLoggingConfigReference{
//   	QueryLoggingConfigArn: jsii.String("queryLoggingConfigArn"),
//   }
//
type QueryLoggingConfigReference struct {
	// The Arn of the QueryLoggingConfig resource.
	QueryLoggingConfigArn *string `field:"required" json:"queryLoggingConfigArn" yaml:"queryLoggingConfigArn"`
}

