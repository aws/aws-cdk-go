package awsdatazone


// The details of the credentials required to access an Amazon Redshift cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftCredentialConfigurationProperty := &RedshiftCredentialConfigurationProperty{
//   	SecretManagerArn: jsii.String("secretManagerArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftcredentialconfiguration.html
//
type CfnDataSource_RedshiftCredentialConfigurationProperty struct {
	// The ARN of a secret manager for an Amazon Redshift cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftcredentialconfiguration.html#cfn-datazone-datasource-redshiftcredentialconfiguration-secretmanagerarn
	//
	SecretManagerArn *string `field:"required" json:"secretManagerArn" yaml:"secretManagerArn"`
}

