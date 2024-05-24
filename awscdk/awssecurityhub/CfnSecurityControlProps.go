package awssecurityhub


// Properties for defining a `CfnSecurityControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityControlProps := &CfnSecurityControlProps{
//   	Parameters: map[string]interface{}{
//   		"parametersKey": &ParameterConfigurationProperty{
//   			"valueType": jsii.String("valueType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	LastUpdateReason: jsii.String("lastUpdateReason"),
//   	SecurityControlArn: jsii.String("securityControlArn"),
//   	SecurityControlId: jsii.String("securityControlId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html
//
type CfnSecurityControlProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html#cfn-securityhub-securitycontrol-parameters
	//
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// The most recent reason for updating the customizable properties of a security control.
	//
	// This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html#cfn-securityhub-securitycontrol-lastupdatereason
	//
	LastUpdateReason *string `field:"optional" json:"lastUpdateReason" yaml:"lastUpdateReason"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html#cfn-securityhub-securitycontrol-securitycontrolarn
	//
	SecurityControlArn *string `field:"optional" json:"securityControlArn" yaml:"securityControlArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html#cfn-securityhub-securitycontrol-securitycontrolid
	//
	SecurityControlId *string `field:"optional" json:"securityControlId" yaml:"securityControlId"`
}

