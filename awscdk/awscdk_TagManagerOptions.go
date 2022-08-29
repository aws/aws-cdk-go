// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options to configure TagManager behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   tagManagerOptions := &tagManagerOptions{
//   	tagPropertyName: jsii.String("tagPropertyName"),
//   }
//
type TagManagerOptions struct {
	// The name of the property in CloudFormation for these tags.
	//
	// Normally this is `tags`, but Cognito UserPool uses UserPoolTags.
	TagPropertyName *string `field:"optional" json:"tagPropertyName" yaml:"tagPropertyName"`
}

