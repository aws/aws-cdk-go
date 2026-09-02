package awselementalinference


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceConfigurationProperty := &DataSourceConfigurationProperty{
//   	FixtureId: jsii.String("fixtureId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-datasourceconfiguration.html
//
type CfnFeed_DataSourceConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elementalinference-feed-datasourceconfiguration.html#cfn-elementalinference-feed-datasourceconfiguration-fixtureid
	//
	FixtureId *string `field:"required" json:"fixtureId" yaml:"fixtureId"`
}

