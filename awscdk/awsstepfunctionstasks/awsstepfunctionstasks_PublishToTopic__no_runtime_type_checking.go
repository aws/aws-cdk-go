//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PublishToTopic) validateBindParameters(_task awsstepfunctions.Task) error {
	return nil
}

func validateNewPublishToTopicParameters(topic awssns.ITopic, props *PublishToTopicProps) error {
	return nil
}

