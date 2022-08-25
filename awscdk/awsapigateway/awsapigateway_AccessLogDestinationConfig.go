package awsapigateway


// Options when binding a log destination to a RestApi Stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogDestinationConfig := &accessLogDestinationConfig{
//   	destinationArn: jsii.String("destinationArn"),
//   }
//
type AccessLogDestinationConfig struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
}

