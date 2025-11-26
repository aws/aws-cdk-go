package previewawsomicsmixins


// Contains information about a source code repository that hosts the workflow definition files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   definitionRepositoryProperty := &DefinitionRepositoryProperty{
//   	ConnectionArn: jsii.String("connectionArn"),
//   	ExcludeFilePatterns: []*string{
//   		jsii.String("excludeFilePatterns"),
//   	},
//   	FullRepositoryId: jsii.String("fullRepositoryId"),
//   	SourceReference: &SourceReferenceProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-definitionrepository.html
//
type CfnWorkflowVersionPropsMixin_DefinitionRepositoryProperty struct {
	// The Amazon Resource Name (ARN) of the connection to the source code repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-definitionrepository.html#cfn-omics-workflowversion-definitionrepository-connectionarn
	//
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
	// A list of file patterns to exclude when retrieving the workflow definition from the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-definitionrepository.html#cfn-omics-workflowversion-definitionrepository-excludefilepatterns
	//
	ExcludeFilePatterns *[]*string `field:"optional" json:"excludeFilePatterns" yaml:"excludeFilePatterns"`
	// The full repository identifier, including the repository owner and name.
	//
	// For example, 'repository-owner/repository-name'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-definitionrepository.html#cfn-omics-workflowversion-definitionrepository-fullrepositoryid
	//
	FullRepositoryId *string `field:"optional" json:"fullRepositoryId" yaml:"fullRepositoryId"`
	// The source reference for the repository, such as a branch name, tag, or commit ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-definitionrepository.html#cfn-omics-workflowversion-definitionrepository-sourcereference
	//
	SourceReference interface{} `field:"optional" json:"sourceReference" yaml:"sourceReference"`
}

