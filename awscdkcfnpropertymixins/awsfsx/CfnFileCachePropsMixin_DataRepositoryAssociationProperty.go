package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataRepositoryAssociationProperty := &DataRepositoryAssociationProperty{
//   	DataRepositoryPath: jsii.String("dataRepositoryPath"),
//   	FileCachePath: jsii.String("fileCachePath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-datarepositoryassociation.html
//
type CfnFileCachePropsMixin_DataRepositoryAssociationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-datarepositoryassociation.html#cfn-fsx-filecache-datarepositoryassociation-datarepositorypath
	//
	DataRepositoryPath *string `field:"optional" json:"dataRepositoryPath" yaml:"dataRepositoryPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-datarepositoryassociation.html#cfn-fsx-filecache-datarepositoryassociation-filecachepath
	//
	FileCachePath *string `field:"optional" json:"fileCachePath" yaml:"fileCachePath"`
}

