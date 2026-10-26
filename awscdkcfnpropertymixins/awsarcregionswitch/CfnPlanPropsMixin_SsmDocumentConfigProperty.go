package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ssmDocumentConfigProperty := &SsmDocumentConfigProperty{
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	DocumentArn: jsii.String("documentArn"),
//   	DocumentVersion: jsii.String("documentVersion"),
//   	ExternalId: jsii.String("externalId"),
//   	Parameters: map[string][]*string{
//   		"parametersKey": []*string{
//   			jsii.String("parameters"),
//   		},
//   	},
//   	Targets: []interface{}{
//   		&SsmTargetProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html
//
type CfnPlanPropsMixin_SsmDocumentConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-documentarn
	//
	DocumentArn *string `field:"optional" json:"documentArn" yaml:"documentArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-documentversion
	//
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmdocumentconfig.html#cfn-arcregionswitch-plan-ssmdocumentconfig-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

