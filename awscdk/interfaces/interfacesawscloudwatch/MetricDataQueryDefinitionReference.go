package interfacesawscloudwatch


// A reference to a MetricDataQueryDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDataQueryDefinitionReference := &MetricDataQueryDefinitionReference{
//   	MetricDataQueryDefinitionArn: jsii.String("metricDataQueryDefinitionArn"),
//   }
//
type MetricDataQueryDefinitionReference struct {
	// The Arn of the MetricDataQueryDefinition resource.
	MetricDataQueryDefinitionArn *string `field:"required" json:"metricDataQueryDefinitionArn" yaml:"metricDataQueryDefinitionArn"`
}

