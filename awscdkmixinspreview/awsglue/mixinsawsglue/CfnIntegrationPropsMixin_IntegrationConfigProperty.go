package mixinsawsglue


// The configuration settings for the integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   integrationConfigProperty := &IntegrationConfigProperty{
//   	ContinuousSync: jsii.Boolean(false),
//   	RefreshInterval: jsii.String("refreshInterval"),
//   	SourceProperties: map[string]*string{
//   		"sourcePropertiesKey": jsii.String("sourceProperties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integration-integrationconfig.html
//
type CfnIntegrationPropsMixin_IntegrationConfigProperty struct {
	// Enables continuous synchronization for on-demand data extractions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integration-integrationconfig.html#cfn-glue-integration-integrationconfig-continuoussync
	//
	ContinuousSync interface{} `field:"optional" json:"continuousSync" yaml:"continuousSync"`
	// Specifies the frequency at which CDC (Change Data Capture) pulls or incremental loads should occur.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integration-integrationconfig.html#cfn-glue-integration-integrationconfig-refreshinterval
	//
	RefreshInterval *string `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
	// A collection of key-value pairs that specify additional properties for the integration source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-integration-integrationconfig.html#cfn-glue-integration-integrationconfig-sourceproperties
	//
	SourceProperties interface{} `field:"optional" json:"sourceProperties" yaml:"sourceProperties"`
}

