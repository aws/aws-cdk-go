package awsiottwinmaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   errorProperty := &errorProperty{
//   	code: jsii.String("code"),
//   	message: jsii.String("message"),
//   }
//
type CfnEntity_ErrorProperty struct {
	// `CfnEntity.ErrorProperty.Code`.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// `CfnEntity.ErrorProperty.Message`.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

