package awscloudfront


// The configuration for a distribution tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tenantConfigProperty := &TenantConfigProperty{
//   	ParameterDefinitions: []interface{}{
//   		&ParameterDefinitionProperty{
//   			Definition: &DefinitionProperty{
//   				StringSchema: &StringSchemaProperty{
//   					Required: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					Comment: jsii.String("comment"),
//   					DefaultValue: jsii.String("defaultValue"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-tenantconfig.html
//
type CfnDistribution_TenantConfigProperty struct {
	// The parameters that you specify for a distribution tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-tenantconfig.html#cfn-cloudfront-distribution-tenantconfig-parameterdefinitions
	//
	ParameterDefinitions interface{} `field:"optional" json:"parameterDefinitions" yaml:"parameterDefinitions"`
}

