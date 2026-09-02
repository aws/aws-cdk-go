package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tagPropagationConfigurationProperty := &TagPropagationConfigurationProperty{
//   	DestinationRoleArn: jsii.String("destinationRoleArn"),
//   	TagConflictResolutionStrategy: jsii.String("tagConflictResolutionStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration.html
//
type CfnOrganizationCentralizationRulePropsMixin_TagPropagationConfigurationProperty struct {
	// The ARN of the destination account IAM role used for tag propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration-destinationrolearn
	//
	DestinationRoleArn *string `field:"optional" json:"destinationRoleArn" yaml:"destinationRoleArn"`
	// The strategy to resolve tag conflicts during propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration-tagconflictresolutionstrategy
	//
	TagConflictResolutionStrategy *string `field:"optional" json:"tagConflictResolutionStrategy" yaml:"tagConflictResolutionStrategy"`
}

