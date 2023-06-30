// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SqlUserState extends com.pulumi.resources.ResourceArgs {

    public static final SqlUserState Empty = new SqlUserState();

    @Import(name="clusterId")
    private @Nullable Output<String> clusterId;

    public Optional<Output<String>> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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

    private SqlUserState() {}

    private SqlUserState(SqlUserState $) {
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.password = $.password;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SqlUserState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SqlUserState $;

        public Builder() {
            $ = new SqlUserState();
        }

        public Builder(SqlUserState defaults) {
            $ = new SqlUserState(Objects.requireNonNull(defaults));
        }

        public Builder clusterId(@Nullable Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

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

        public SqlUserState build() {
            return $;
        }
    }

}