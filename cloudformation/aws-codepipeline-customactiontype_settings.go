package cloudformation

// AWSCodePipelineCustomActionType_Settings AWS CloudFormation Resource (AWS::CodePipeline::CustomActionType.Settings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html
type AWSCodePipelineCustomActionType_Settings struct {

	// EntityUrlTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-entityurltemplate
	EntityUrlTemplate *stringIntrinsic `json:"EntityUrlTemplate,omitempty"`

	// ExecutionUrlTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-executionurltemplate
	ExecutionUrlTemplate *stringIntrinsic `json:"ExecutionUrlTemplate,omitempty"`

	// RevisionUrlTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-revisionurltemplate
	RevisionUrlTemplate *stringIntrinsic `json:"RevisionUrlTemplate,omitempty"`

	// ThirdPartyConfigurationUrl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-thirdpartyconfigurationurl
	ThirdPartyConfigurationUrl *stringIntrinsic `json:"ThirdPartyConfigurationUrl,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelineCustomActionType_Settings) AWSCloudFormationType() string {
	return "AWS::CodePipeline::CustomActionType.Settings"
}
