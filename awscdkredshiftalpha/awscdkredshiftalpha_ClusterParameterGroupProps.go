// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha


// Properties for a parameter group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import redshift_alpha "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//
//   clusterParameterGroupProps := &clusterParameterGroupProps{
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
// Experimental.
type ClusterParameterGroupProps struct {
	// The parameters in this parameter group.
	// Experimental.
	Parameters *map[string]*string `field:"required" json:"parameters" yaml:"parameters"`
	// Description for this parameter group.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

