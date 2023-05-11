package awsiottwinmaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   errorProperty := &ErrorProperty{
//   	Code: jsii.String("code"),
//   	Message: jsii.String("message"),
//   }
//
type CfnComponentType_ErrorProperty struct {
	// `CfnComponentType.ErrorProperty.Code`.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// `CfnComponentType.ErrorProperty.Message`.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

