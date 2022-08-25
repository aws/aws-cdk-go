package awsec2


// Describes a resource statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceStatementRequestProperty := &resourceStatementRequestProperty{
//   	resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	resourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   }
//
type CfnNetworkInsightsAccessScope_ResourceStatementRequestProperty struct {
	// The resources.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The resource types.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

