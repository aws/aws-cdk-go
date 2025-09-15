package awskafkaconnect


// A reference to a CustomPlugin resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPluginReference := &CustomPluginReference{
//   	CustomPluginArn: jsii.String("customPluginArn"),
//   }
//
type CustomPluginReference struct {
	// The CustomPluginArn of the CustomPlugin resource.
	CustomPluginArn *string `field:"required" json:"customPluginArn" yaml:"customPluginArn"`
}

