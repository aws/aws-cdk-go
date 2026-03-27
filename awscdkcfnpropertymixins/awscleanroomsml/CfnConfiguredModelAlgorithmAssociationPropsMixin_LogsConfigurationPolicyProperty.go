package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   logsConfigurationPolicyProperty := &LogsConfigurationPolicyProperty{
//   	AllowedAccountIds: []*string{
//   		jsii.String("allowedAccountIds"),
//   	},
//   	FilterPattern: jsii.String("filterPattern"),
//   	LogRedactionConfiguration: &LogRedactionConfigurationProperty{
//   		CustomEntityConfig: &CustomEntityConfigProperty{
//   			CustomDataIdentifiers: []*string{
//   				jsii.String("customDataIdentifiers"),
//   			},
//   		},
//   		EntitiesToRedact: []*string{
//   			jsii.String("entitiesToRedact"),
//   		},
//   	},
//   	LogType: jsii.String("logType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy.html
//
type CfnConfiguredModelAlgorithmAssociationPropsMixin_LogsConfigurationPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy-allowedaccountids
	//
	AllowedAccountIds *[]*string `field:"optional" json:"allowedAccountIds" yaml:"allowedAccountIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy-filterpattern
	//
	FilterPattern *string `field:"optional" json:"filterPattern" yaml:"filterPattern"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy-logredactionconfiguration
	//
	LogRedactionConfiguration interface{} `field:"optional" json:"logRedactionConfiguration" yaml:"logRedactionConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-logsconfigurationpolicy-logtype
	//
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
}

