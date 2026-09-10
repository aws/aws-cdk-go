package awselementalinference


// Identifies the fixture whose event data Elemental Inference maps onto the clipping metadata for an output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	FixtureId: jsii.String("fixtureId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-datasourceconfiguration.html
//
type CfnFeedPropsMixin_DataSourceConfigurationProperty struct {
	// The ID of the fixture whose event data you want Elemental Inference to map onto this clipping output.
	//
	// To obtain this ID, use the SearchFixtures operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-datasourceconfiguration.html#cfn-elementalinference-feed-datasourceconfiguration-fixtureid
	//
	FixtureId *string `field:"optional" json:"fixtureId" yaml:"fixtureId"`
}

