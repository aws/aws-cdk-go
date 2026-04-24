package awscdk


// Resource violating a specific rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   policyViolatingResourceBeta1 := &PolicyViolatingResourceBeta1{
//   	Locations: []*string{
//   		jsii.String("locations"),
//   	},
//   	ResourceLogicalId: jsii.String("resourceLogicalId"),
//   	TemplatePath: jsii.String("templatePath"),
//   }
//
// Deprecated: Use `PolicyViolatingResource` instead.
type PolicyViolatingResourceBeta1 struct {
	// The locations in the CloudFormation template that pose the violations.
	// Deprecated: Use `PolicyViolatingResource` instead.
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// The logical ID of the resource in the CloudFormation template.
	// Deprecated: Use `PolicyViolatingResource` instead.
	ResourceLogicalId *string `field:"required" json:"resourceLogicalId" yaml:"resourceLogicalId"`
	// The path to the CloudFormation template that contains this resource.
	// Deprecated: Use `PolicyViolatingResource` instead.
	TemplatePath *string `field:"required" json:"templatePath" yaml:"templatePath"`
}

