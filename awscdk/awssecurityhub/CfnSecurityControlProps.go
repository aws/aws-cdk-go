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
//
//   			// the properties below are optional
//   			"value": &ParameterValueProperty{
//   				"boolean": jsii.Boolean(false),
//   				"double": jsii.Number(123),
//   				"enum": jsii.String("enum"),
//   				"enumList": []*string{
//   					jsii.String("enumList"),
//   				},
//   				"integer": jsii.Number(123),
//   				"integerList": []interface{}{
//   					jsii.Number(123),
//   				},
//   				"string": jsii.String("string"),
//   				"stringList": []*string{
//   					jsii.String("stringList"),
//   				},
//   			},
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
	// An object that identifies the name of a control parameter, its current value, and whether it has been customized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html#cfn-securityhub-securitycontrol-parameters
	//
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// The most recent reason for updating the customizable properties of a security control.
	//
	// This differs from the `UpdateReason` field of the [`BatchUpdateStandardsControlAssociations`](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_BatchUpdateStandardsControlAssociations.html) API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html#cfn-securityhub-securitycontrol-lastupdatereason
	//
	LastUpdateReason *string `field:"optional" json:"lastUpdateReason" yaml:"lastUpdateReason"`
	// The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1` . This parameter doesn't mention a specific standard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html#cfn-securityhub-securitycontrol-securitycontrolarn
	//
	SecurityControlArn *string `field:"optional" json:"securityControlArn" yaml:"securityControlArn"`
	// The unique identifier of a security control across standards.
	//
	// Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-securitycontrol.html#cfn-securityhub-securitycontrol-securitycontrolid
	//
	SecurityControlId *string `field:"optional" json:"securityControlId" yaml:"securityControlId"`
}

