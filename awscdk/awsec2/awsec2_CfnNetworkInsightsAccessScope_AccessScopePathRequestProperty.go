package awsec2


// Describes a path.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessScopePathRequestProperty := &accessScopePathRequestProperty{
//   	destination: &pathStatementRequestProperty{
//   		packetHeaderStatement: &packetHeaderStatementRequestProperty{
//   			destinationAddresses: []*string{
//   				jsii.String("destinationAddresses"),
//   			},
//   			destinationPorts: []*string{
//   				jsii.String("destinationPorts"),
//   			},
//   			destinationPrefixLists: []*string{
//   				jsii.String("destinationPrefixLists"),
//   			},
//   			protocols: []*string{
//   				jsii.String("protocols"),
//   			},
//   			sourceAddresses: []*string{
//   				jsii.String("sourceAddresses"),
//   			},
//   			sourcePorts: []*string{
//   				jsii.String("sourcePorts"),
//   			},
//   			sourcePrefixLists: []*string{
//   				jsii.String("sourcePrefixLists"),
//   			},
//   		},
//   		resourceStatement: &resourceStatementRequestProperty{
//   			resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			resourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   		},
//   	},
//   	source: &pathStatementRequestProperty{
//   		packetHeaderStatement: &packetHeaderStatementRequestProperty{
//   			destinationAddresses: []*string{
//   				jsii.String("destinationAddresses"),
//   			},
//   			destinationPorts: []*string{
//   				jsii.String("destinationPorts"),
//   			},
//   			destinationPrefixLists: []*string{
//   				jsii.String("destinationPrefixLists"),
//   			},
//   			protocols: []*string{
//   				jsii.String("protocols"),
//   			},
//   			sourceAddresses: []*string{
//   				jsii.String("sourceAddresses"),
//   			},
//   			sourcePorts: []*string{
//   				jsii.String("sourcePorts"),
//   			},
//   			sourcePrefixLists: []*string{
//   				jsii.String("sourcePrefixLists"),
//   			},
//   		},
//   		resourceStatement: &resourceStatementRequestProperty{
//   			resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			resourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   		},
//   	},
//   	throughResources: []interface{}{
//   		&throughResourcesStatementRequestProperty{
//   			resourceStatement: &resourceStatementRequestProperty{
//   				resources: []*string{
//   					jsii.String("resources"),
//   				},
//   				resourceTypes: []*string{
//   					jsii.String("resourceTypes"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnNetworkInsightsAccessScope_AccessScopePathRequestProperty struct {
	// The destination.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The source.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The through resources.
	ThroughResources interface{} `field:"optional" json:"throughResources" yaml:"throughResources"`
}

