package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   interceptorInputConfigurationProperty := &InterceptorInputConfigurationProperty{
//   	PassRequestHeaders: jsii.Boolean(false),
//   	PayloadFilter: &InterceptorPayloadFilterProperty{
//   		Exclude: []interface{}{
//   			&InterceptorPayloadExclusionSelectorProperty{
//   				Field: jsii.String("field"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorinputconfiguration.html
//
type CfnGatewayPropsMixin_InterceptorInputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorinputconfiguration.html#cfn-bedrockagentcore-gateway-interceptorinputconfiguration-passrequestheaders
	//
	PassRequestHeaders interface{} `field:"optional" json:"passRequestHeaders" yaml:"passRequestHeaders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorinputconfiguration.html#cfn-bedrockagentcore-gateway-interceptorinputconfiguration-payloadfilter
	//
	PayloadFilter interface{} `field:"optional" json:"payloadFilter" yaml:"payloadFilter"`
}

