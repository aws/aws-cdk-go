package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ssmSendCommandConfigurationProperty := &SsmSendCommandConfigurationProperty{
//   	RegionDocuments: map[string]interface{}{
//   		"regionDocumentsKey": &SsmDocumentConfigProperty{
//   			"crossAccountRole": jsii.String("crossAccountRole"),
//   			"documentArn": jsii.String("documentArn"),
//   			"documentVersion": jsii.String("documentVersion"),
//   			"externalId": jsii.String("externalId"),
//   			"parameters": map[string][]*string{
//   				"parametersKey": []*string{
//   					jsii.String("parameters"),
//   				},
//   			},
//   			"targets": []interface{}{
//   				&SsmTargetProperty{
//   					"key": jsii.String("key"),
//   					"values": []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RetryIntervalMinutes: jsii.Number(123),
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &SsmSendCommandUngracefulProperty{
//   		Behavior: jsii.String("behavior"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html
//
type CfnPlanPropsMixin_SsmSendCommandConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html#cfn-arcregionswitch-plan-ssmsendcommandconfiguration-regiondocuments
	//
	RegionDocuments interface{} `field:"optional" json:"regionDocuments" yaml:"regionDocuments"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html#cfn-arcregionswitch-plan-ssmsendcommandconfiguration-retryintervalminutes
	//
	RetryIntervalMinutes *float64 `field:"optional" json:"retryIntervalMinutes" yaml:"retryIntervalMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html#cfn-arcregionswitch-plan-ssmsendcommandconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ssmsendcommandconfiguration.html#cfn-arcregionswitch-plan-ssmsendcommandconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

