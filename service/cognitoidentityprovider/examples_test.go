// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cognitoidentityprovider_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCognitoIdentityProvider_AddCustomAttributes() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AddCustomAttributesInput{
		CustomAttributes: []*cognitoidentityprovider.SchemaAttributeType{ // Required
			{ // Required
				AttributeDataType:      aws.String("AttributeDataType"),
				DeveloperOnlyAttribute: aws.Bool(true),
				Mutable:                aws.Bool(true),
				Name:                   aws.String("CustomAttributeNameType"),
				NumberAttributeConstraints: &cognitoidentityprovider.NumberAttributeConstraintsType{
					MaxValue: aws.String("StringType"),
					MinValue: aws.String("StringType"),
				},
				Required: aws.Bool(true),
				StringAttributeConstraints: &cognitoidentityprovider.StringAttributeConstraintsType{
					MaxLength: aws.String("StringType"),
					MinLength: aws.String("StringType"),
				},
			},
			// More values...
		},
		UserPoolId: aws.String("UserPoolIdType"), // Required
	}
	resp, err := svc.AddCustomAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_AdminConfirmSignUp() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
		Username:   aws.String("UsernameType"),   // Required
	}
	resp, err := svc.AdminConfirmSignUp(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_AdminDeleteUser() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
		Username:   aws.String("UsernameType"),   // Required
	}
	resp, err := svc.AdminDeleteUser(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_AdminDeleteUserAttributes() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AdminDeleteUserAttributesInput{
		UserAttributeNames: []*string{ // Required
			aws.String("AttributeNameType"), // Required
			// More values...
		},
		UserPoolId: aws.String("UserPoolIdType"), // Required
		Username:   aws.String("UsernameType"),   // Required
	}
	resp, err := svc.AdminDeleteUserAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_AdminDisableUser() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AdminDisableUserInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
		Username:   aws.String("UsernameType"),   // Required
	}
	resp, err := svc.AdminDisableUser(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_AdminEnableUser() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AdminEnableUserInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
		Username:   aws.String("UsernameType"),   // Required
	}
	resp, err := svc.AdminEnableUser(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_AdminGetUser() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
		Username:   aws.String("UsernameType"),   // Required
	}
	resp, err := svc.AdminGetUser(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_AdminResetUserPassword() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AdminResetUserPasswordInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
		Username:   aws.String("UsernameType"),   // Required
	}
	resp, err := svc.AdminResetUserPassword(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_AdminSetUserSettings() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AdminSetUserSettingsInput{
		MFAOptions: []*cognitoidentityprovider.MFAOptionType{ // Required
			{ // Required
				AttributeName:  aws.String("AttributeNameType"),
				DeliveryMedium: aws.String("DeliveryMediumType"),
			},
			// More values...
		},
		UserPoolId: aws.String("UserPoolIdType"), // Required
		Username:   aws.String("UsernameType"),   // Required
	}
	resp, err := svc.AdminSetUserSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_AdminUpdateUserAttributes() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserAttributes: []*cognitoidentityprovider.AttributeType{ // Required
			{ // Required
				Name:  aws.String("AttributeNameType"), // Required
				Value: aws.String("AttributeValueType"),
			},
			// More values...
		},
		UserPoolId: aws.String("UserPoolIdType"), // Required
		Username:   aws.String("UsernameType"),   // Required
	}
	resp, err := svc.AdminUpdateUserAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_ChangePassword() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.ChangePasswordInput{
		PreviousPassword: aws.String("PasswordType"), // Required
		ProposedPassword: aws.String("PasswordType"), // Required
		AccessToken:      aws.String("TokenModelType"),
	}
	resp, err := svc.ChangePassword(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_ConfirmForgotPassword() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         aws.String("ClientIdType"),         // Required
		ConfirmationCode: aws.String("ConfirmationCodeType"), // Required
		Password:         aws.String("PasswordType"),         // Required
		Username:         aws.String("UsernameType"),         // Required
		SecretHash:       aws.String("SecretHashType"),
	}
	resp, err := svc.ConfirmForgotPassword(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_ConfirmSignUp() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:           aws.String("ClientIdType"),         // Required
		ConfirmationCode:   aws.String("ConfirmationCodeType"), // Required
		Username:           aws.String("UsernameType"),         // Required
		ForceAliasCreation: aws.Bool(true),
		SecretHash:         aws.String("SecretHashType"),
	}
	resp, err := svc.ConfirmSignUp(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_CreateUserPool() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.CreateUserPoolInput{
		PoolName: aws.String("UserPoolNameType"), // Required
		AliasAttributes: []*string{
			aws.String("AliasAttributeType"), // Required
			// More values...
		},
		AutoVerifiedAttributes: []*string{
			aws.String("VerifiedAttributeType"), // Required
			// More values...
		},
		EmailVerificationMessage: aws.String("EmailVerificationMessageType"),
		EmailVerificationSubject: aws.String("EmailVerificationSubjectType"),
		LambdaConfig: &cognitoidentityprovider.LambdaConfigType{
			CustomMessage:      aws.String("ArnType"),
			PostAuthentication: aws.String("ArnType"),
			PostConfirmation:   aws.String("ArnType"),
			PreAuthentication:  aws.String("ArnType"),
			PreSignUp:          aws.String("ArnType"),
		},
		MfaConfiguration: aws.String("UserPoolMfaType"),
		Policies: &cognitoidentityprovider.UserPoolPolicyType{
			PasswordPolicy: &cognitoidentityprovider.PasswordPolicyType{
				MinimumLength:    aws.Int64(1),
				RequireLowercase: aws.Bool(true),
				RequireNumbers:   aws.Bool(true),
				RequireSymbols:   aws.Bool(true),
				RequireUppercase: aws.Bool(true),
			},
		},
		SmsAuthenticationMessage: aws.String("SmsVerificationMessageType"),
		SmsVerificationMessage:   aws.String("SmsVerificationMessageType"),
	}
	resp, err := svc.CreateUserPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_CreateUserPoolClient() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.CreateUserPoolClientInput{
		ClientName:     aws.String("ClientNameType"), // Required
		UserPoolId:     aws.String("UserPoolIdType"), // Required
		GenerateSecret: aws.Bool(true),
	}
	resp, err := svc.CreateUserPoolClient(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_DeleteUser() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.DeleteUserInput{
		AccessToken: aws.String("TokenModelType"),
	}
	resp, err := svc.DeleteUser(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_DeleteUserAttributes() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.DeleteUserAttributesInput{
		UserAttributeNames: []*string{ // Required
			aws.String("AttributeNameType"), // Required
			// More values...
		},
		AccessToken: aws.String("TokenModelType"),
	}
	resp, err := svc.DeleteUserAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_DeleteUserPool() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.DeleteUserPoolInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
	}
	resp, err := svc.DeleteUserPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_DeleteUserPoolClient() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.DeleteUserPoolClientInput{
		ClientId:   aws.String("ClientIdType"),   // Required
		UserPoolId: aws.String("UserPoolIdType"), // Required
	}
	resp, err := svc.DeleteUserPoolClient(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_DescribeUserPool() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.DescribeUserPoolInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
	}
	resp, err := svc.DescribeUserPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_DescribeUserPoolClient() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.DescribeUserPoolClientInput{
		ClientId:   aws.String("ClientIdType"),   // Required
		UserPoolId: aws.String("UserPoolIdType"), // Required
	}
	resp, err := svc.DescribeUserPoolClient(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_ForgotPassword() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.ForgotPasswordInput{
		ClientId:   aws.String("ClientIdType"), // Required
		Username:   aws.String("UsernameType"), // Required
		SecretHash: aws.String("SecretHashType"),
	}
	resp, err := svc.ForgotPassword(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_GetJWKS() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.GetJWKSInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
	}
	resp, err := svc.GetJWKS(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_GetOpenIdConfiguration() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.GetOpenIdConfigurationInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
	}
	resp, err := svc.GetOpenIdConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_GetUser() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String("TokenModelType"),
	}
	resp, err := svc.GetUser(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_GetUserAttributeVerificationCode() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.GetUserAttributeVerificationCodeInput{
		AttributeName: aws.String("AttributeNameType"), // Required
		AccessToken:   aws.String("TokenModelType"),
	}
	resp, err := svc.GetUserAttributeVerificationCode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_ListUserPoolClients() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.ListUserPoolClientsInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationKey"),
	}
	resp, err := svc.ListUserPoolClients(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_ListUserPools() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: aws.Int64(1), // Required
		NextToken:  aws.String("PaginationKeyType"),
	}
	resp, err := svc.ListUserPools(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_ListUsers() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
		AttributesToGet: []*string{
			aws.String("AttributeNameType"), // Required
			// More values...
		},
		Limit:           aws.Int64(1),
		PaginationToken: aws.String("SearchPaginationTokenType"),
		UserStatus:      aws.String("UserStatusType"),
	}
	resp, err := svc.ListUsers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_ResendConfirmationCode() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId:   aws.String("ClientIdType"), // Required
		Username:   aws.String("UsernameType"), // Required
		SecretHash: aws.String("SecretHashType"),
	}
	resp, err := svc.ResendConfirmationCode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_SetUserSettings() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.SetUserSettingsInput{
		AccessToken: aws.String("TokenModelType"), // Required
		MFAOptions: []*cognitoidentityprovider.MFAOptionType{ // Required
			{ // Required
				AttributeName:  aws.String("AttributeNameType"),
				DeliveryMedium: aws.String("DeliveryMediumType"),
			},
			// More values...
		},
	}
	resp, err := svc.SetUserSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_SignUp() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String("ClientIdType"), // Required
		Password:   aws.String("PasswordType"), // Required
		Username:   aws.String("UsernameType"), // Required
		SecretHash: aws.String("SecretHashType"),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{ // Required
				Name:  aws.String("AttributeNameType"), // Required
				Value: aws.String("AttributeValueType"),
			},
			// More values...
		},
		ValidationData: []*cognitoidentityprovider.AttributeType{
			{ // Required
				Name:  aws.String("AttributeNameType"), // Required
				Value: aws.String("AttributeValueType"),
			},
			// More values...
		},
	}
	resp, err := svc.SignUp(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_UpdateUserAttributes() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.UpdateUserAttributesInput{
		UserAttributes: []*cognitoidentityprovider.AttributeType{ // Required
			{ // Required
				Name:  aws.String("AttributeNameType"), // Required
				Value: aws.String("AttributeValueType"),
			},
			// More values...
		},
		AccessToken: aws.String("TokenModelType"),
	}
	resp, err := svc.UpdateUserAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_UpdateUserPool() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.UpdateUserPoolInput{
		UserPoolId: aws.String("UserPoolIdType"), // Required
		AutoVerifiedAttributes: []*string{
			aws.String("VerifiedAttributeType"), // Required
			// More values...
		},
		EmailVerificationMessage: aws.String("EmailVerificationMessageType"),
		EmailVerificationSubject: aws.String("EmailVerificationSubjectType"),
		LambdaConfig: &cognitoidentityprovider.LambdaConfigType{
			CustomMessage:      aws.String("ArnType"),
			PostAuthentication: aws.String("ArnType"),
			PostConfirmation:   aws.String("ArnType"),
			PreAuthentication:  aws.String("ArnType"),
			PreSignUp:          aws.String("ArnType"),
		},
		MfaConfiguration: aws.String("UserPoolMfaType"),
		Policies: &cognitoidentityprovider.UserPoolPolicyType{
			PasswordPolicy: &cognitoidentityprovider.PasswordPolicyType{
				MinimumLength:    aws.Int64(1),
				RequireLowercase: aws.Bool(true),
				RequireNumbers:   aws.Bool(true),
				RequireSymbols:   aws.Bool(true),
				RequireUppercase: aws.Bool(true),
			},
		},
		SmsAuthenticationMessage: aws.String("SmsVerificationMessageType"),
		SmsVerificationMessage:   aws.String("SmsVerificationMessageType"),
	}
	resp, err := svc.UpdateUserPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_UpdateUserPoolClient() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.UpdateUserPoolClientInput{
		ClientId:   aws.String("ClientIdType"),   // Required
		UserPoolId: aws.String("UserPoolIdType"), // Required
		ClientName: aws.String("ClientNameType"),
	}
	resp, err := svc.UpdateUserPoolClient(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentityProvider_VerifyUserAttribute() {
	svc := cognitoidentityprovider.New(session.New())

	params := &cognitoidentityprovider.VerifyUserAttributeInput{
		AttributeName: aws.String("AttributeNameType"),    // Required
		Code:          aws.String("ConfirmationCodeType"), // Required
		AccessToken:   aws.String("TokenModelType"),
	}
	resp, err := svc.VerifyUserAttribute(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}