// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Cockroach
{
    public static class GetPersonUser
    {
        /// <summary>
        /// Information about an individual user.
        /// </summary>
        public static Task<GetPersonUserResult> InvokeAsync(GetPersonUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPersonUserResult>("cockroach:index/getPersonUser:getPersonUser", args ?? new GetPersonUserArgs(), options.WithDefaults());

        /// <summary>
        /// Information about an individual user.
        /// </summary>
        public static Output<GetPersonUserResult> Invoke(GetPersonUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPersonUserResult>("cockroach:index/getPersonUser:getPersonUser", args ?? new GetPersonUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPersonUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Email address used to find the User ID.
        /// </summary>
        [Input("email", required: true)]
        public string Email { get; set; } = null!;

        public GetPersonUserArgs()
        {
        }
        public static new GetPersonUserArgs Empty => new GetPersonUserArgs();
    }

    public sealed class GetPersonUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Email address used to find the User ID.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        public GetPersonUserInvokeArgs()
        {
        }
        public static new GetPersonUserInvokeArgs Empty => new GetPersonUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetPersonUserResult
    {
        /// <summary>
        /// Email address used to find the User ID.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// User ID.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPersonUserResult(
            string email,

            string id)
        {
            Email = email;
            Id = id;
        }
    }
}
