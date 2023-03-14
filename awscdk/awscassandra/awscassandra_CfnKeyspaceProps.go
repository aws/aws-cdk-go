package awscassandra

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnKeyspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKeyspaceProps := &cfnKeyspaceProps{
//   	keyspaceName: jsii.String("keyspaceName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnKeyspaceProps struct {
	// The name of the keyspace to be created.
	//
	// The keyspace name is case sensitive. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the keyspace name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// *Length constraints:* Minimum length of 3. Maximum length of 255.
	//
	// *Pattern:* `^[a-zA-Z0-9][a-zA-Z0-9_]{1,47}$`.
	KeyspaceName *string `field:"optional" json:"keyspaceName" yaml:"keyspaceName"`
	// A list of key-value pair tags to be attached to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

