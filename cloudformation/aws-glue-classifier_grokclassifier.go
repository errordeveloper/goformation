package cloudformation

// AWSGlueClassifier_GrokClassifier AWS CloudFormation Resource (AWS::Glue::Classifier.GrokClassifier)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-grokclassifier.html
type AWSGlueClassifier_GrokClassifier struct {

	// Classification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-grokclassifier.html#cfn-glue-classifier-grokclassifier-classification
	Classification *stringIntrinsic `json:"Classification,omitempty"`

	// CustomPatterns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-grokclassifier.html#cfn-glue-classifier-grokclassifier-custompatterns
	CustomPatterns *stringIntrinsic `json:"CustomPatterns,omitempty"`

	// GrokPattern AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-grokclassifier.html#cfn-glue-classifier-grokclassifier-grokpattern
	GrokPattern *stringIntrinsic `json:"GrokPattern,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-grokclassifier.html#cfn-glue-classifier-grokclassifier-name
	Name *stringIntrinsic `json:"Name,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueClassifier_GrokClassifier) AWSCloudFormationType() string {
	return "AWS::Glue::Classifier.GrokClassifier"
}
