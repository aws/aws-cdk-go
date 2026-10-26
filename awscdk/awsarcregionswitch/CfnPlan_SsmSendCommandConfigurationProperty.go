package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmSendCommandConfigurationProperty := &SsmSendCommandConfigurationProperty{
//   	RegionDocuments: map[string]interface{}{
//   		"regionDocumentsKey": &SsmDocumentConfigProperty{
//   			"documentArn": jsii.String("documentArn"),
//   			"targets": []interface{}{
//   				&SsmTargetProperty{
//   					"key": jsii.String("key"),
//   					"values": []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			"crossAccountRole": jsii.String("crossAccountRole"),
//   			"documentVersion": jsii.String("documentVersion"),
//   			"externalId": jsii.String("externalId"),
//   			"parameters": map[string][]*string{
//   				"parametersKey": []*string{
//   					jsii.String("parameters"),
//   				},
//   			},
//   		},
//   	},
//   	RetryIntervalMinutes: jsii.Number(123),
//
//   	// the properties below are optional
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &SsmSendCommandUngracefulProperty{
//   		Behavior: jsii.String("behavior"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html
//
type CfnPlan_SsmSendCommandConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html#cfn-arcregionswitch-plan-ssmsendcommandconfiguration-regiondocuments
	//
	RegionDocuments interface{} `field:"required" json:"regionDocuments" yaml:"regionDocuments"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html#cfn-arcregionswitch-plan-ssmsendcommandconfiguration-retryintervalminutes
	//
	RetryIntervalMinutes *float64 `field:"required" json:"retryIntervalMinutes" yaml:"retryIntervalMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html#cfn-arcregionswitch-plan-ssmsendcommandconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html#cfn-arcregionswitch-plan-ssmsendcommandconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

