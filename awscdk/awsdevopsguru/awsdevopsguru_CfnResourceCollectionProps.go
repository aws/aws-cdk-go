package awsdevopsguru


// Properties for defining a `CfnResourceCollection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceCollectionProps := &cfnResourceCollectionProps{
//   	resourceCollectionFilter: &resourceCollectionFilterProperty{
//   		cloudFormation: &cloudFormationCollectionFilterProperty{
//   			stackNames: []*string{
//   				jsii.String("stackNames"),
//   			},
//   		},
//   		tags: []tagCollectionProperty{
//   			&tagCollectionProperty{
//   				appBoundaryKey: jsii.String("appBoundaryKey"),
//   				tagValues: []*string{
//   					jsii.String("tagValues"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnResourceCollectionProps struct {
	// Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
	ResourceCollectionFilter interface{} `field:"required" json:"resourceCollectionFilter" yaml:"resourceCollectionFilter"`
}

