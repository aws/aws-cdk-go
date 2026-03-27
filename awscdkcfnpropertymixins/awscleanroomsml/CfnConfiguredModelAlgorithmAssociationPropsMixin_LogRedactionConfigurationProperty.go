package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   logRedactionConfigurationProperty := &LogRedactionConfigurationProperty{
//   	CustomEntityConfig: &CustomEntityConfigProperty{
//   		CustomDataIdentifiers: []*string{
//   			jsii.String("customDataIdentifiers"),
//   		},
//   	},
//   	EntitiesToRedact: []*string{
//   		jsii.String("entitiesToRedact"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration.html
//
type CfnConfiguredModelAlgorithmAssociationPropsMixin_LogRedactionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration-customentityconfig
	//
	CustomEntityConfig interface{} `field:"optional" json:"customEntityConfig" yaml:"customEntityConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-logredactionconfiguration-entitiestoredact
	//
	EntitiesToRedact *[]*string `field:"optional" json:"entitiesToRedact" yaml:"entitiesToRedact"`
}

