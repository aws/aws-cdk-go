package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   customEntityConfigProperty := &CustomEntityConfigProperty{
//   	CustomDataIdentifiers: []*string{
//   		jsii.String("customDataIdentifiers"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-customentityconfig.html
//
type CfnConfiguredModelAlgorithmAssociationPropsMixin_CustomEntityConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-customentityconfig.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-customentityconfig-customdataidentifiers
	//
	CustomDataIdentifiers *[]*string `field:"optional" json:"customDataIdentifiers" yaml:"customDataIdentifiers"`
}

