// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sts provides the client and types for making API
// requests to AWS STS.
//
// The AWS Security Token Service (STS) is a web service that enables you to
// request temporary, limited-privilege credentials for AWS Identity and Access
// Management (IAM) users or for users that you authenticate (federated users).
// This guide provides descriptions of the STS API. For more detailed information
// about using this service, go to Temporary Security Credentials (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html).
//
// As an alternative to using the API, you can use one of the AWS SDKs, which
// consist of libraries and sample code for various programming languages and
// platforms (Java, Ruby, .NET, iOS, Android, etc.). The SDKs provide a convenient
// way to create programmatic access to STS. For example, the SDKs take care
// of cryptographically signing requests, managing errors, and retrying requests
// automatically. For information about the AWS SDKs, including how to download
// and install them, see the Tools for Amazon Web Services page (http://aws.amazon.com/tools/).
//
// For information about setting up signatures and authorization through the
// API, go to Signing AWS API Requests (https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html)
// in the AWS General Reference. For general information about the Query API,
// go to Making Query Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html)
// in Using IAM. For information about using security tokens with other AWS
// products, go to AWS Services That Work with IAM (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html)
// in the IAM User Guide.
//
// If you're new to AWS and need additional technical information about a specific
// AWS product, you can find the product's technical documentation at http://aws.amazon.com/documentation/
// (http://aws.amazon.com/documentation/).
//
// Endpoints
//
// By default, AWS Security Token Service (STS) is available as a global service,
// and all AWS STS requests go to a single endpoint at https://sts.amazonaws.com.
// Global requests map to the US East (N. Virginia) region. AWS recommends using
// Regional AWS STS endpoints instead of the global endpoint to reduce latency,
// build in redundancy, and increase session token validity. For more information,
// see Managing AWS STS in an AWS Region (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the IAM User Guide.
//
// Most AWS Regions are enabled for operations in all AWS services by default.
// Those Regions are automatically activated for use with AWS STS. Some Regions,
// such as Asia Pacific (Hong Kong), must be manually enabled. To learn more
// about enabling and disabling AWS Regions, see Managing AWS Regions (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html)
// in the AWS General Reference. When you enable these AWS Regions, they are
// automatically activated for use with AWS STS. You cannot activate the STS
// endpoint for a Region that is disabled. Tokens that are valid in all AWS
// Regions are longer than tokens that are valid in Regions that are enabled
// by default. Changing this setting might affect existing systems where you
// temporarily store tokens. For more information, see Managing Global Endpoint
// Session Tokens (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html#sts-regions-manage-tokens)
// in the IAM User Guide.
//
// After you activate a Region for use with AWS STS, you can direct AWS STS
// API calls to that Region. AWS STS recommends that you use both the setRegion
// and setEndpoint methods to make calls to a Regional endpoint. You can use
// the setRegion method alone for manually enabled Regions, such as Asia Pacific
// (Hong Kong). In this case, the calls are directed to the STS Regional endpoint.
// However, if you use the setRegion method alone for Regions enabled by default,
// the calls are directed to the global endpoint of https://sts.amazonaws.com.
//
// To view the list of AWS STS endpoints and whether they are active by default,
// see Writing Code to Use AWS STS Regions (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html#id_credentials_temp_enable-regions_writing_code)
// in the IAM User Guide.
//
// Recording API requests
//
// STS supports AWS CloudTrail, which is a service that records AWS calls for
// your AWS account and delivers log files to an Amazon S3 bucket. By using
// information collected by CloudTrail, you can determine what requests were
// successfully made to STS, who made the request, when it was made, and so
// on.
//
// If you activate AWS STS endpoints in Regions other than the default global
// endpoint, then you must also turn on CloudTrail logging in those Regions.
// This is necessary to record any AWS STS API calls that are made in those
// Regions. For more information, see Turning On CloudTrail in Additional Regions
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/aggregating_logs_regions_turn_on_ct.html)
// in the AWS CloudTrail User Guide.
//
// AWS Security Token Service (STS) is a global service with a single endpoint
// at https://sts.amazonaws.com. Calls to this endpoint are logged as calls
// to a global service. However, because this endpoint is physically located
// in the US East (N. Virginia) Region, your logs list us-east-1 as the event
// Region. CloudTrail does not write these logs to the US East (Ohio) Region
// unless you choose to include global service logs in that Region. CloudTrail
// writes calls to all Regional endpoints to their respective Regions. For example,
// calls to sts.us-east-2.amazonaws.com are published to the US East (Ohio)
// Region and calls to sts.eu-central-1.amazonaws.com are published to the EU
// (Frankfurt) Region.
//
// To learn more about CloudTrail, including how to turn it on and find your
// log files, see the AWS CloudTrail User Guide (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/what_is_cloud_trail_top_level.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15 for more information on this service.
//
// See sts package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/sts/
//
// Using the Client
//
// To use AWS STS with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS STS client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/sts/#New
package sts
