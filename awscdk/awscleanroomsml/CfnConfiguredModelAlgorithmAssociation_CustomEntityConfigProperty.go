package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customEntityConfigProperty := &CustomEntityConfigProperty{
//   	CustomDataIdentifiers: []*string{
//   		jsii.String("customDataIdentifiers"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-customentityconfig.html
//
type CfnConfiguredModelAlgorithmAssociation_CustomEntityConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-customentityconfig.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-customentityconfig-customdataidentifiers
	//
	CustomDataIdentifiers *[]*string `field:"required" json:"customDataIdentifiers" yaml:"customDataIdentifiers"`
}

