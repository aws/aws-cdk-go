package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsAuthConfigurationProperty := &TlsAuthConfigurationProperty{
//   	TrustStore: &TrustStoreProperty{
//   		CaContent: jsii.String("caContent"),
//
//   		// the properties below are optional
//   		CrlContent: jsii.String("crlContent"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-tlsauthconfiguration.html
//
type CfnMailManagerIngressPoint_TlsAuthConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-tlsauthconfiguration.html#cfn-ses-mailmanageringresspoint-tlsauthconfiguration-truststore
	//
	TrustStore interface{} `field:"required" json:"trustStore" yaml:"trustStore"`
}

