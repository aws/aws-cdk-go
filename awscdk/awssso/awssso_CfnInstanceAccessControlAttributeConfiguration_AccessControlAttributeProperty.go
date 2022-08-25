package awssso


// These are AWS SSO identity store attributes that you can configure for use in attributes-based access control (ABAC).
//
// You can create permissions policies that determine who can access your AWS resources based upon the configured attribute values. When you enable ABAC and specify `AccessControlAttributes` , AWS SSO passes the attribute values of the authenticated user into IAM for use in policy evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlAttributeProperty := &accessControlAttributeProperty{
//   	key: jsii.String("key"),
//   	value: &accessControlAttributeValueProperty{
//   		source: []*string{
//   			jsii.String("source"),
//   		},
//   	},
//   }
//
type CfnInstanceAccessControlAttributeConfiguration_AccessControlAttributeProperty struct {
	// The name of the attribute associated with your identities in your identity source.
	//
	// This is used to map a specified attribute in your identity source with an attribute in AWS SSO .
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value used for mapping a specified attribute to an identity source.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

