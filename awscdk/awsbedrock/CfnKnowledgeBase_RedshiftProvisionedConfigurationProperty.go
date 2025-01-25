package awsbedrock


// Configurations for provisioned Redshift query engine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftProvisionedConfigurationProperty := &RedshiftProvisionedConfigurationProperty{
//   	AuthConfiguration: &RedshiftProvisionedAuthConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		DatabaseUser: jsii.String("databaseUser"),
//   		UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   	},
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedconfiguration.html
//
type CfnKnowledgeBase_RedshiftProvisionedConfigurationProperty struct {
	// Configurations for Redshift query engine provisioned auth setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedconfiguration-authconfiguration
	//
	AuthConfiguration interface{} `field:"required" json:"authConfiguration" yaml:"authConfiguration"`
	// Redshift cluster identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedconfiguration-clusteridentifier
	//
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

