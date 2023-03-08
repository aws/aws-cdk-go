package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyTextConfigProperty := &clarifyTextConfigProperty{
//   	granularity: jsii.String("granularity"),
//   	language: jsii.String("language"),
//   }
//
type CfnEndpointConfig_ClarifyTextConfigProperty struct {
	// `CfnEndpointConfig.ClarifyTextConfigProperty.Granularity`.
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// `CfnEndpointConfig.ClarifyTextConfigProperty.Language`.
	Language *string `field:"required" json:"language" yaml:"language"`
}

