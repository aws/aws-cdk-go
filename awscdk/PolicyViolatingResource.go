package awscdk


// Resource violating a specific rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   policyViolatingResource := &PolicyViolatingResource{
//   	Locations: []*string{
//   		jsii.String("locations"),
//   	},
//   	ResourceLogicalId: jsii.String("resourceLogicalId"),
//   	TemplatePath: jsii.String("templatePath"),
//   }
//
type PolicyViolatingResource struct {
	// The locations in the CloudFormation template that pose the violations.
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// The logical ID of the resource in the CloudFormation template.
	ResourceLogicalId *string `field:"required" json:"resourceLogicalId" yaml:"resourceLogicalId"`
	// The path to the CloudFormation template that contains this resource.
	TemplatePath *string `field:"required" json:"templatePath" yaml:"templatePath"`
}

