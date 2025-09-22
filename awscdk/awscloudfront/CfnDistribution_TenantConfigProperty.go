package awscloudfront


// > This field only supports multi-tenant distributions.
//
// You can't specify this field for standard distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-config-options.html#unsupported-saas) in the *Amazon CloudFront Developer Guide* .
//
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

