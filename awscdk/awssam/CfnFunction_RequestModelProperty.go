package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestModelProperty := &RequestModelProperty{
//   	Model: jsii.String("model"),
//
//   	// the properties below are optional
//   	Required: jsii.Boolean(false),
//   	ValidateBody: jsii.Boolean(false),
//   	ValidateParameters: jsii.Boolean(false),
//   }
//
type CfnFunction_RequestModelProperty struct {
	// `CfnFunction.RequestModelProperty.Model`.
	Model *string `field:"required" json:"model" yaml:"model"`
	// `CfnFunction.RequestModelProperty.Required`.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// `CfnFunction.RequestModelProperty.ValidateBody`.
	ValidateBody interface{} `field:"optional" json:"validateBody" yaml:"validateBody"`
	// `CfnFunction.RequestModelProperty.ValidateParameters`.
	ValidateParameters interface{} `field:"optional" json:"validateParameters" yaml:"validateParameters"`
}

