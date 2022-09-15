// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appconfigdata provides the client and types for making API
// requests to AWS AppConfig Data.
//
// AppConfig Data provides the data plane APIs your application uses to retrieve
// configuration data. Here's how it works:
//
// Your application retrieves configuration data by first establishing a configuration
// session using the AppConfig Data StartConfigurationSession API action. Your
// session's client then makes periodic calls to GetLatestConfiguration to check
// for and retrieve the latest data available.
//
// When calling StartConfigurationSession, your code sends the following information:
//
//   - Identifiers (ID or name) of an AppConfig application, environment, and
//     configuration profile that the session tracks.
//
//   - (Optional) The minimum amount of time the session's client must wait
//     between calls to GetLatestConfiguration.
//
// In response, AppConfig provides an InitialConfigurationToken to be given
// to the session's client and used the first time it calls GetLatestConfiguration
// for that session.
//
// When calling GetLatestConfiguration, your client code sends the most recent
// ConfigurationToken value it has and receives in response:
//
//   - NextPollConfigurationToken: the ConfigurationToken value to use on the
//     next call to GetLatestConfiguration.
//
//   - NextPollIntervalInSeconds: the duration the client should wait before
//     making its next call to GetLatestConfiguration. This duration may vary
//     over the course of the session, so it should be used instead of the value
//     sent on the StartConfigurationSession call.
//
//   - The configuration: the latest data intended for the session. This may
//     be empty if the client already has the latest version of the configuration.
//
// For more information and to view example CLI commands that show how to retrieve
// a configuration using the AppConfig Data StartConfigurationSession and GetLatestConfiguration
// API actions, see Receiving the configuration (http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration)
// in the AppConfig User Guide.
//
// See https://docs.aws.amazon.com/goto/WebAPI/appconfigdata-2021-11-11 for more information on this service.
//
// See appconfigdata package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/appconfigdata/
//
// # Using the Client
//
// To contact AWS AppConfig Data with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS AppConfig Data client AppConfigData for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/appconfigdata/#New
package appconfigdata