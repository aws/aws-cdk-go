package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchEnvironmentVariableProperty := &batchEnvironmentVariableProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnPipe_BatchEnvironmentVariableProperty struct {
	// `CfnPipe.BatchEnvironmentVariableProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnPipe.BatchEnvironmentVariableProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

