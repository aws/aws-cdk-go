package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestParameterProperty := &requestParameterProperty{
//   	caching: jsii.Boolean(false),
//   	required: jsii.Boolean(false),
//   }
//
type CfnFunction_RequestParameterProperty struct {
	// `CfnFunction.RequestParameterProperty.Caching`.
	Caching interface{} `field:"optional" json:"caching" yaml:"caching"`
	// `CfnFunction.RequestParameterProperty.Required`.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

