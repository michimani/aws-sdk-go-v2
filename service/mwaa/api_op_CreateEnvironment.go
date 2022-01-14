// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.
func (c *Client) CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) {
	if params == nil {
		params = &CreateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironment", params, optFns, c.addOperationCreateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This section contains the Amazon Managed Workflows for Apache Airflow (MWAA) API
// reference documentation to create an environment. For more information, see Get
// started with Amazon Managed Workflows for Apache Airflow
// (https://docs.aws.amazon.com/mwaa/latest/userguide/get-started.html).
type CreateEnvironmentInput struct {

	// The relative path to the DAGs folder on your Amazon S3 bucket. For example,
	// dags. To learn more, see Adding or updating DAGs
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-folder.html).
	//
	// This member is required.
	DagS3Path *string

	// The Amazon Resource Name (ARN) of the execution role for your environment. An
	// execution role is an Amazon Web Services Identity and Access Management (IAM)
	// role that grants MWAA permission to access Amazon Web Services services and
	// resources used by your environment. For example,
	// arn:aws:iam::123456789:role/my-execution-role. To learn more, see Amazon MWAA
	// Execution role
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html).
	//
	// This member is required.
	ExecutionRoleArn *string

	// The name of the Amazon MWAA environment. For example, MyMWAAEnvironment.
	//
	// This member is required.
	Name *string

	// The VPC networking components used to secure and enable network traffic between
	// the Amazon Web Services resources for your environment. To learn more, see About
	// networking on Amazon MWAA
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html).
	//
	// This member is required.
	NetworkConfiguration *types.NetworkConfiguration

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and
	// supporting files are stored. For example,
	// arn:aws:s3:::my-airflow-bucket-unique-name. To learn more, see Create an Amazon
	// S3 bucket for Amazon MWAA
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-s3-bucket.html).
	//
	// This member is required.
	SourceBucketArn *string

	// A list of key-value pairs containing the Apache Airflow configuration options
	// you want to attach to your environment. To learn more, see Apache Airflow
	// configuration options
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html).
	AirflowConfigurationOptions map[string]string

	// The Apache Airflow version for your environment. If no value is specified,
	// defaults to the latest version. Valid values: 1.10.12, 2.0.2. To learn more, see
	// Apache Airflow versions on Amazon Managed Workflows for Apache Airflow (MWAA)
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/airflow-versions.html).
	AirflowVersion *string

	// The environment class type. Valid values: mw1.small, mw1.medium, mw1.large. To
	// learn more, see Amazon MWAA environment class
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/environment-class.html).
	EnvironmentClass *string

	// The Amazon Web Services Key Management Service (KMS) key to encrypt the data in
	// your environment. You can use an Amazon Web Services owned CMK, or a Customer
	// managed CMK (advanced). To learn more, see Create an Amazon MWAA environment
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/create-environment.html).
	KmsKey *string

	// Defines the Apache Airflow logs to send to CloudWatch Logs.
	LoggingConfiguration *types.LoggingConfigurationInput

	// The maximum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify in the
	// MaxWorkers field. For example, 20. When there are no more tasks running, and no
	// more in the queue, MWAA disposes of the extra workers leaving the one worker
	// that is included with your environment, or the number you specify in MinWorkers.
	MaxWorkers *int32

	// The minimum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify in the
	// MaxWorkers field. When there are no more tasks running, and no more in the
	// queue, MWAA disposes of the extra workers leaving the worker count you specify
	// in the MinWorkers field. For example, 2.
	MinWorkers *int32

	// The version of the plugins.zip file on your Amazon S3 bucket. A version must be
	// specified each time a plugins.zip file is updated. To learn more, see How S3
	// Versioning works
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html).
	PluginsS3ObjectVersion *string

	// The relative path to the plugins.zip file on your Amazon S3 bucket. For example,
	// plugins.zip. If specified, then the plugins.zip version is required. To learn
	// more, see Installing custom plugins
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html).
	PluginsS3Path *string

	// The version of the requirements.txt file on your Amazon S3 bucket. A version
	// must be specified each time a requirements.txt file is updated. To learn more,
	// see How S3 Versioning works
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html).
	RequirementsS3ObjectVersion *string

	// The relative path to the requirements.txt file on your Amazon S3 bucket. For
	// example, requirements.txt. If specified, then a file version is required. To
	// learn more, see Installing Python dependencies
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html).
	RequirementsS3Path *string

	// The number of Apache Airflow schedulers to run in your environment. Valid
	// values:
	//
	// * v2.0.2 - Accepts between 2 to 5. Defaults to 2.
	//
	// * v1.10.12 - Accepts
	// 1.
	Schedulers *int32

	// The key-value tag pairs you want to associate to your environment. For example,
	// "Environment": "Staging". To learn more, see Tagging Amazon Web Services
	// resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags map[string]string

	// The Apache Airflow Web server access mode. To learn more, see Apache Airflow
	// access modes
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-networking.html).
	WebserverAccessMode types.WebserverAccessMode

	// The day and time of the week in Coordinated Universal Time (UTC) 24-hour
	// standard time to start weekly maintenance updates of your environment in the
	// following format: DAY:HH:MM. For example: TUE:03:30. You can specify a start
	// time in 30 minute increments only.
	WeeklyMaintenanceWindowStart *string

	noSmithyDocumentSerde
}

type CreateEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) returned in the response for the environment.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opCreateEnvironmentMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironment(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opCreateEnvironmentMiddleware struct {
}

func (*endpointPrefix_opCreateEnvironmentMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateEnvironmentMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateEnvironmentMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateEnvironmentMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "airflow",
		OperationName: "CreateEnvironment",
	}
}
