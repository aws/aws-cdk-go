package awsekslegacy

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//   var vpc vpc
//
//   clusterAttributes := &ClusterAttributes{
//   	ClusterArn: jsii.String("clusterArn"),
//   	ClusterCertificateAuthorityData: jsii.String("clusterCertificateAuthorityData"),
//   	ClusterEndpoint: jsii.String("clusterEndpoint"),
//   	ClusterName: jsii.String("clusterName"),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	Vpc: vpc,
//   }
//
// Experimental.
type ClusterAttributes struct {
	// The unique ARN assigned to the service by AWS in the form of arn:aws:eks:.
	// Experimental.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// The certificate-authority-data for your cluster.
	// Experimental.
	ClusterCertificateAuthorityData *string `field:"required" json:"clusterCertificateAuthorityData" yaml:"clusterCertificateAuthorityData"`
	// The API Server endpoint URL.
	// Experimental.
	ClusterEndpoint *string `field:"required" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// The physical name of the Cluster.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The security groups associated with this cluster.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The VPC in which this Cluster was created.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

