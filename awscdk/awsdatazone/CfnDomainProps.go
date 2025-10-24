package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainProps := &CfnDomainProps{
//   	DomainExecutionRole: jsii.String("domainExecutionRole"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DomainVersion: jsii.String("domainVersion"),
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	ServiceRole: jsii.String("serviceRole"),
//   	SingleSignOn: &SingleSignOnProperty{
//   		IdcInstanceArn: jsii.String("idcInstanceArn"),
//   		Type: jsii.String("type"),
//   		UserAssignment: jsii.String("userAssignment"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domain.html
//
type CfnDomainProps struct {
	// The domain execution role that is created when an Amazon DataZone domain is created.
	//
	// The domain execution role is created in the AWS account that houses the Amazon DataZone domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domain.html#cfn-datazone-domain-domainexecutionrole
	//
	DomainExecutionRole *string `field:"required" json:"domainExecutionRole" yaml:"domainExecutionRole"`
	// The name of the Amazon DataZone domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domain.html#cfn-datazone-domain-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the Amazon DataZone domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domain.html#cfn-datazone-domain-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The domain version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domain.html#cfn-datazone-domain-domainversion
	//
	DomainVersion *string `field:"optional" json:"domainVersion" yaml:"domainVersion"`
	// The identifier of the AWS Key Management Service (KMS) key that is used to encrypt the Amazon DataZone domain, metadata, and reporting data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domain.html#cfn-datazone-domain-kmskeyidentifier
	//
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The service role of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domain.html#cfn-datazone-domain-servicerole
	//
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The single sign-on details in Amazon DataZone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domain.html#cfn-datazone-domain-singlesignon
	//
	SingleSignOn interface{} `field:"optional" json:"singleSignOn" yaml:"singleSignOn"`
	// The tags specified for the Amazon DataZone domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-domain.html#cfn-datazone-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

