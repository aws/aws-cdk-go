package awsrds


// A reference to a DBProxy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBProxyReference := &DBProxyReference{
//   	DbProxyArn: jsii.String("dbProxyArn"),
//   	DbProxyName: jsii.String("dbProxyName"),
//   }
//
type DBProxyReference struct {
	// The ARN of the DBProxy resource.
	DbProxyArn *string `field:"required" json:"dbProxyArn" yaml:"dbProxyArn"`
	// The DBProxyName of the DBProxy resource.
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
}

