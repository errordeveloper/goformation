package main_test

import (
	"encoding/json"

	"github.com/awslabs/goformation/cloudformation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource", func() {

	Context("with a resource that contains a mix of primitive, custom and polymorphic properties", func() {

		Context("described as Go structs", func() {

			Context("with a simple primitive used for a polymorphic property", func() {

				codeuri := cloudformation.String("s3://bucket/key")
				resource := &cloudformation.AWSServerlessFunction{
					Runtime: cloudformation.String("nodejs6.10"),
					CodeUri: codeuri,
				}

				expected := []byte(`{"Type":"AWS::Serverless::Function","Properties":{"CodeUri":"s3://bucket/key","Runtime":"nodejs6.10"}}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a custom type used for a polymorphic property", func() {

				resource := &cloudformation.AWSServerlessFunction{
					Runtime: cloudformation.String("nodejs6.10"),
					CodeUri: &cloudformation.AWSServerlessFunction_S3Location{
						Bucket:  cloudformation.String("test-bucket"),
						Key:     cloudformation.String("test-key"),
						Version: cloudformation.Integer(123),
					},
				}

				expected := []byte(`{"Type":"AWS::Serverless::Function","Properties":{"CodeUri":{"Bucket":"test-bucket","Key":"test-key","Version":123},"Runtime":"nodejs6.10"}}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

	})

	Context("with a custom property type resource", func() {

		Context("described as Go structs", func() {

			Context("with a list type", func() {

				subproperty := &cloudformation.AWSServerlessFunction_S3Event{
					Bucket: cloudformation.String("my-bucket"),
					Events: cloudformation.MakeStringSlice("s3:ObjectCreated:*", "s3:ObjectRemoved:*"),
				}

				expected := []byte(`{"Bucket":"my-bucket","Events":["s3:ObjectCreated:*","s3:ObjectRemoved:*"]}`)

				result, err := json.Marshal(subproperty)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a primitive type", func() {

				event := cloudformation.String("s3:ObjectCreated:*")
				subproperty := &cloudformation.AWSServerlessFunction_S3Event{
					Bucket: cloudformation.String("my-bucket"),
					Events: event,
				}

				expected := []byte(`{"Bucket":"my-bucket","Events":"s3:ObjectCreated:*"}`)

				result, err := json.Marshal(subproperty)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})
	})
})
