package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoEventProperty := &CognitoEventProperty{
//   	Trigger: jsii.String("trigger"),
//   	UserPool: jsii.String("userPool"),
//   }
//
type CfnFunction_CognitoEventProperty struct {
	// `CfnFunction.CognitoEventProperty.Trigger`.
	Trigger interface{} `field:"required" json:"trigger" yaml:"trigger"`
	// `CfnFunction.CognitoEventProperty.UserPool`.
	UserPool *string `field:"required" json:"userPool" yaml:"userPool"`
}

