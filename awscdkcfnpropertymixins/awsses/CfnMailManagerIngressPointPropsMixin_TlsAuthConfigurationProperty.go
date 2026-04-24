package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tlsAuthConfigurationProperty := &TlsAuthConfigurationProperty{
//   	TrustStore: &TrustStoreProperty{
//   		CaContent: jsii.String("caContent"),
//   		CrlContent: jsii.String("crlContent"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-tlsauthconfiguration.html
//
type CfnMailManagerIngressPointPropsMixin_TlsAuthConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-tlsauthconfiguration.html#cfn-ses-mailmanageringresspoint-tlsauthconfiguration-truststore
	//
	TrustStore interface{} `field:"optional" json:"trustStore" yaml:"trustStore"`
}

