package awsec2


// Describes a through resource statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   throughResourcesStatementRequestProperty := &throughResourcesStatementRequestProperty{
//   	resourceStatement: &resourceStatementRequestProperty{
//   		resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		resourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   	},
//   }
//
type CfnNetworkInsightsAccessScope_ThroughResourcesStatementRequestProperty struct {
	// The resource statement.
	ResourceStatement interface{} `field:"optional" json:"resourceStatement" yaml:"resourceStatement"`
}

