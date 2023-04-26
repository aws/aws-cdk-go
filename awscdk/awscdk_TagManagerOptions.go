// An experiment to bundle the entire CDK into a single module
package awscdk


// Options to configure TagManager behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   tagManagerOptions := &TagManagerOptions{
//   	TagPropertyName: jsii.String("tagPropertyName"),
//   }
//
// Experimental.
type TagManagerOptions struct {
	// The name of the property in CloudFormation for these tags.
	//
	// Normally this is `tags`, but Cognito UserPool uses UserPoolTags.
	// Experimental.
	TagPropertyName *string `field:"optional" json:"tagPropertyName" yaml:"tagPropertyName"`
}

