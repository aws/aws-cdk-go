package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagPropagationConfigurationProperty := &TagPropagationConfigurationProperty{
//   	DestinationRoleArn: jsii.String("destinationRoleArn"),
//
//   	// the properties below are optional
//   	TagConflictResolutionStrategy: jsii.String("tagConflictResolutionStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration.html
//
type CfnOrganizationCentralizationRule_TagPropagationConfigurationProperty struct {
	// The ARN of the destination account IAM role used for tag propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration-destinationrolearn
	//
	DestinationRoleArn *string `field:"required" json:"destinationRoleArn" yaml:"destinationRoleArn"`
	// The strategy to resolve tag conflicts during propagation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-tagpropagationconfiguration-tagconflictresolutionstrategy
	//
	TagConflictResolutionStrategy *string `field:"optional" json:"tagConflictResolutionStrategy" yaml:"tagConflictResolutionStrategy"`
}

