package awslogs


// This processor adds new key-value pairs to the log event.
//
// For more information about this processor including examples, see addKeys in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorDeleteKeysProperty := &ProcessorDeleteKeysProperty{
//   	WithKeys: []string{
//   		jsii.String("withKeys"),
//   	},
//   }
//
type ProcessorDeleteKeysProperty struct {
	// A list of keys to delete.
	WithKeys *[]*string `field:"required" json:"withKeys" yaml:"withKeys"`
}

