package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmDocumentConfigProperty := &SsmDocumentConfigProperty{
//   	DocumentArn: jsii.String("documentArn"),
//   	Targets: []interface{}{
//   		&SsmTargetProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	DocumentVersion: jsii.String("documentVersion"),
//   	ExternalId: jsii.String("externalId"),
//   	Parameters: map[string][]*string{
//   		"parametersKey": []*string{
//   			jsii.String("parameters"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html
//
type CfnPlan_SsmDocumentConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-documentarn
	//
	DocumentArn *string `field:"required" json:"documentArn" yaml:"documentArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-targets
	//
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-documentversion
	//
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

