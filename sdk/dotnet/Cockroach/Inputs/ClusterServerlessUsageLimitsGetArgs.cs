// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Cockroach.Inputs
{

    public sealed class ClusterServerlessUsageLimitsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("requestUnitLimit", required: true)]
        public Input<int> RequestUnitLimit { get; set; } = null!;

        [Input("storageMibLimit", required: true)]
        public Input<int> StorageMibLimit { get; set; } = null!;

        public ClusterServerlessUsageLimitsGetArgs()
        {
        }
        public static new ClusterServerlessUsageLimitsGetArgs Empty => new ClusterServerlessUsageLimitsGetArgs();
    }
}
