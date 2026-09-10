package awscdk


// Information about a single stack that is being validated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   policyValidationStack := &PolicyValidationStack{
//   	StackConstructPath: jsii.String("stackConstructPath"),
//   	TemplatePath: jsii.String("templatePath"),
//
//   	// the properties below are optional
//   	AccountId: jsii.String("accountId"),
//   	Region: jsii.String("region"),
//   }
//
type PolicyValidationStack struct {
	// The Stack's construct path.
	StackConstructPath *string `field:"required" json:"stackConstructPath" yaml:"stackConstructPath"`
	// The path to the template file on disk.
	TemplatePath *string `field:"required" json:"templatePath" yaml:"templatePath"`
	// The account ID for this stack, if known.
	// Default: - the account ID is unknown.
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The region for this stack, if known.
	// Default: - the region is unknown.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

