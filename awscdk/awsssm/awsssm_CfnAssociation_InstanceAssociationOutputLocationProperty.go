package awsssm


// `InstanceAssociationOutputLocation` is a property of the [AWS::SSM::Association](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html) resource that specifies an Amazon S3 bucket where you want to store the results of this association request.
//
// For the minimal permissions required to enable Amazon S3 output for an association, see [Creating associations](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-state-assoc.html) in the *Systems Manager User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceAssociationOutputLocationProperty := &instanceAssociationOutputLocationProperty{
//   	s3Location: &s3OutputLocationProperty{
//   		outputS3BucketName: jsii.String("outputS3BucketName"),
//   		outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   		outputS3Region: jsii.String("outputS3Region"),
//   	},
//   }
//
type CfnAssociation_InstanceAssociationOutputLocationProperty struct {
	// `S3OutputLocation` is a property of the [InstanceAssociationOutputLocation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html) property that specifies an Amazon S3 bucket where you want to store the results of this request.
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

