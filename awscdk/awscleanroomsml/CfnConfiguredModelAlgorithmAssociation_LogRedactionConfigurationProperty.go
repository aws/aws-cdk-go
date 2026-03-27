package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logRedactionConfigurationProperty := &LogRedactionConfigurationProperty{
//   	EntitiesToRedact: []*string{
//   		jsii.String("entitiesToRedact"),
//   	},
//
//   	// the properties below are optional
//   	CustomEntityConfig: &CustomEntityConfigProperty{
//   		CustomDataIdentifiers: []*string{
//   			jsii.String("customDataIdentifiers"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration.html
//
type CfnConfiguredModelAlgorithmAssociation_LogRedactionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration-entitiestoredact
	//
	EntitiesToRedact *[]*string `field:"required" json:"entitiesToRedact" yaml:"entitiesToRedact"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration-customentityconfig
	//
	CustomEntityConfig interface{} `field:"optional" json:"customEntityConfig" yaml:"customEntityConfig"`
}

