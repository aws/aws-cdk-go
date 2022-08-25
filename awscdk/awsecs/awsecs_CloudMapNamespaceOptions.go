package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
)

// The options for creating an AWS Cloud Map namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   cloudMapNamespaceOptions := &cloudMapNamespaceOptions{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	type: awscdk.Aws_servicediscovery.namespaceType_HTTP,
//   	vpc: vpc,
//   }
//
type CloudMapNamespaceOptions struct {
	// The name of the namespace, such as example.com.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of CloudMap Namespace to create.
	Type awsservicediscovery.NamespaceType `field:"optional" json:"type" yaml:"type"`
	// The VPC to associate the namespace with.
	//
	// This property is required for private DNS namespaces.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

