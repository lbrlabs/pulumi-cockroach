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
    public static class GetOrganization
    {
        /// <summary>
        /// Information about the organization associated with the user's API key.
        /// </summary>
        public static Task<GetOrganizationResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationResult>("cockroach:index/getOrganization:getOrganization", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Information about the organization associated with the user's API key.
        /// </summary>
        public static Output<GetOrganizationResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationResult>("cockroach:index/getOrganization:getOrganization", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetOrganizationResult
    {
        /// <summary>
        /// Indicates when the organization was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Organization ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A short ID used by CockroachDB Support.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// Name of the organization.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetOrganizationResult(
            string createdAt,

            string id,

            string label,

            string name)
        {
            CreatedAt = createdAt;
            Id = id;
            Label = label;
            Name = name;
        }
    }
}
