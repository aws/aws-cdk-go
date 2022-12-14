package awsopensearchserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samlConfigOptionsProperty := &samlConfigOptionsProperty{
//   	metadata: jsii.String("metadata"),
//
//   	// the properties below are optional
//   	groupAttribute: jsii.String("groupAttribute"),
//   	sessionTimeout: jsii.Number(123),
//   	userAttribute: jsii.String("userAttribute"),
//   }
//
type CfnSecurityConfig_SamlConfigOptionsProperty struct {
	// `CfnSecurityConfig.SamlConfigOptionsProperty.Metadata`.
	Metadata *string `field:"required" json:"metadata" yaml:"metadata"`
	// `CfnSecurityConfig.SamlConfigOptionsProperty.GroupAttribute`.
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// `CfnSecurityConfig.SamlConfigOptionsProperty.SessionTimeout`.
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// `CfnSecurityConfig.SamlConfigOptionsProperty.UserAttribute`.
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

