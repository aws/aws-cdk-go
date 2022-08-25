package awscodepipeline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   actionConfig := &actionConfig{
//   	configuration: configuration,
//   }
//
type ActionConfig struct {
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
}

