package interfacesawsgroundstation


// A reference to a Config resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configReference := &ConfigReference{
//   	ConfigArn: jsii.String("configArn"),
//   }
//
type ConfigReference struct {
	// The Arn of the Config resource.
	ConfigArn *string `field:"required" json:"configArn" yaml:"configArn"`
}

