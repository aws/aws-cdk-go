package awsresiliencehubv2


// An input source for the service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSourceDefinitionProperty := &InputSourceDefinitionProperty{
//   	ResourceConfiguration: &ResourceConfigurationProperty{
//   		CfnStackArn: jsii.String("cfnStackArn"),
//   		DesignFileS3Url: jsii.String("designFileS3Url"),
//   		Eks: &EksSourceProperty{
//   			ClusterArn: jsii.String("clusterArn"),
//   			Namespaces: []*string{
//   				jsii.String("namespaces"),
//   			},
//   		},
//   		ResourceTags: []interface{}{
//   			&ResourceTagProperty{
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		TfStateFileUrl: jsii.String("tfStateFileUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-inputsourcedefinition.html
//
type CfnService_InputSourceDefinitionProperty struct {
	// Resource configuration for an input source.
	//
	// Provide exactly one field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-inputsourcedefinition.html#cfn-resiliencehubv2-service-inputsourcedefinition-resourceconfiguration
	//
	ResourceConfiguration interface{} `field:"required" json:"resourceConfiguration" yaml:"resourceConfiguration"`
}

