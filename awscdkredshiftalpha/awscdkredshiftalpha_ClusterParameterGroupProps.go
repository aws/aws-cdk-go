// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha


// Properties for a parameter group.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//
//
//   params := awscdkredshiftalpha.NewClusterParameterGroup(this, jsii.String("Params"), &clusterParameterGroupProps{
//   	description: jsii.String("desc"),
//   	parameters: map[string]*string{
//   		"require_ssl": jsii.String("true"),
//   	},
//   })
//
//   params.addParameter(jsii.String("enable_user_activity_logging"), jsii.String("true"))
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

