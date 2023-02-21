package awsiot


// Parameters to define a mitigation action that adds a blank policy to restrict permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replaceDefaultPolicyVersionParamsProperty := &replaceDefaultPolicyVersionParamsProperty{
//   	templateName: jsii.String("templateName"),
//   }
//
type CfnMitigationAction_ReplaceDefaultPolicyVersionParamsProperty struct {
	// The name of the template to be applied.
	//
	// The only supported value is `BLANK_POLICY` .
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
}

