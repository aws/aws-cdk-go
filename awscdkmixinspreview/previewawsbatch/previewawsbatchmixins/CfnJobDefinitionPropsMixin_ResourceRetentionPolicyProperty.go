package previewawsbatchmixins


// Specifies the resource retention policy settings for a job definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceRetentionPolicyProperty := &ResourceRetentionPolicyProperty{
//   	SkipDeregisterOnUpdate: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resourceretentionpolicy.html
//
type CfnJobDefinitionPropsMixin_ResourceRetentionPolicyProperty struct {
	// Specifies whether the previous revision of the job definition is retained in an active status after UPDATE events for the resource.
	//
	// The default value is `false` . When the property is set to `false` , the previous revision of the job definition is de-registered after a new revision is created. When the property is set to `true` , the previous revision of the job definition is not de-registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resourceretentionpolicy.html#cfn-batch-jobdefinition-resourceretentionpolicy-skipderegisteronupdate
	//
	// Default: - false.
	//
	SkipDeregisterOnUpdate interface{} `field:"optional" json:"skipDeregisterOnUpdate" yaml:"skipDeregisterOnUpdate"`
}

