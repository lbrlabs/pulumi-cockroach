// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SqlUserArgs extends com.pulumi.resources.ResourceArgs {

    public static final SqlUserArgs Empty = new SqlUserArgs();

    @Import(name="clusterId", required=true)
    private Output<String> clusterId;

    public Output<String> clusterId() {
        return this.clusterId;
    }

    /**
     * SQL user name.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return SQL user name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * If provided, this field sets the password of the SQL user when created. If omitted, a random password is generated, but
     * not saved to Terraform state. The password must be changed via the CockroachDB cloud console.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return If provided, this field sets the password of the SQL user when created. If omitted, a random password is generated, but
     * not saved to Terraform state. The password must be changed via the CockroachDB cloud console.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    private SqlUserArgs() {}

    private SqlUserArgs(SqlUserArgs $) {
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.password = $.password;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SqlUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SqlUserArgs $;

        public Builder() {
            $ = new SqlUserArgs();
        }

        public Builder(SqlUserArgs defaults) {
            $ = new SqlUserArgs(Objects.requireNonNull(defaults));
        }

        public Builder clusterId(Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param name SQL user name.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name SQL user name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password If provided, this field sets the password of the SQL user when created. If omitted, a random password is generated, but
         * not saved to Terraform state. The password must be changed via the CockroachDB cloud console.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password If provided, this field sets the password of the SQL user when created. If omitted, a random password is generated, but
         * not saved to Terraform state. The password must be changed via the CockroachDB cloud console.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        public SqlUserArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
