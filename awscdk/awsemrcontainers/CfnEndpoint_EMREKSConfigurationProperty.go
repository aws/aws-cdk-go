package awsemrcontainers


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eMREKSConfigurationProperty_ EMREKSConfigurationProperty
//
//   eMREKSConfigurationProperty := &EMREKSConfigurationProperty{
//   	Classification: jsii.String("classification"),
//
//   	// the properties below are optional
//   	Configurations: []interface{}{
//   		eMREKSConfigurationProperty_,
//   	},
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-emreksconfiguration.html
//
type CfnEndpoint_EMREKSConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-emreksconfiguration.html#cfn-emrcontainers-endpoint-emreksconfiguration-classification
	//
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-emreksconfiguration.html#cfn-emrcontainers-endpoint-emreksconfiguration-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-emreksconfiguration.html#cfn-emrcontainers-endpoint-emreksconfiguration-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

