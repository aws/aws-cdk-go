package awsses


// The configuration of the ingress endpoint resource.
//
// > This data type is a UNION, so only one of the following members can be specified when used or returned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressPointConfigurationProperty := &IngressPointConfigurationProperty{
//   	SecretArn: jsii.String("secretArn"),
//   	SmtpPassword: jsii.String("smtpPassword"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-ingresspointconfiguration.html
//
type CfnMailManagerIngressPoint_IngressPointConfigurationProperty struct {
	// The SecretsManager::Secret ARN of the ingress endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-ingresspointconfiguration.html#cfn-ses-mailmanageringresspoint-ingresspointconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The password of the ingress endpoint resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-ingresspointconfiguration.html#cfn-ses-mailmanageringresspoint-ingresspointconfiguration-smtppassword
	//
	SmtpPassword *string `field:"optional" json:"smtpPassword" yaml:"smtpPassword"`
}

