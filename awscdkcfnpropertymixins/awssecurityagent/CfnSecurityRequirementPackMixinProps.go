package awssecurityagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSecurityRequirementPackPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSecurityRequirementPackMixinProps := &CfnSecurityRequirementPackMixinProps{
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	SecurityRequirements: []interface{}{
//   		&SecurityRequirementProperty{
//   			Description: jsii.String("description"),
//   			Domain: jsii.String("domain"),
//   			Evaluation: jsii.String("evaluation"),
//   			Name: jsii.String("name"),
//   			Remediation: jsii.String("remediation"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-securityrequirementpack.html
//
type CfnSecurityRequirementPackMixinProps struct {
	// Description of the pack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-securityrequirementpack.html#cfn-securityagent-securityrequirementpack-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// KMS key for client-side encryption of pack contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-securityrequirementpack.html#cfn-securityagent-securityrequirementpack-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Name of the security requirement pack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-securityrequirementpack.html#cfn-securityagent-securityrequirementpack-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Security requirements within this pack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-securityrequirementpack.html#cfn-securityagent-securityrequirementpack-securityrequirements
	//
	SecurityRequirements interface{} `field:"optional" json:"securityRequirements" yaml:"securityRequirements"`
	// Whether the pack is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-securityrequirementpack.html#cfn-securityagent-securityrequirementpack-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Tags for the security requirement pack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-securityrequirementpack.html#cfn-securityagent-securityrequirementpack-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

