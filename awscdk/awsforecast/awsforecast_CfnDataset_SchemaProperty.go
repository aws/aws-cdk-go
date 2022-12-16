package awsforecast


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaProperty := &schemaProperty{
//   	attributes: []interface{}{
//   		&attributesItemsProperty{
//   			attributeName: jsii.String("attributeName"),
//   			attributeType: jsii.String("attributeType"),
//   		},
//   	},
//   }
//
type CfnDataset_SchemaProperty struct {
	// `CfnDataset.SchemaProperty.Attributes`.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
}

