// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetConnectionStringConnectionParams {
    private String database;
    private String host;
    private String password;
    private String port;
    private String username;

    private GetConnectionStringConnectionParams() {}
    public String database() {
        return this.database;
    }
    public String host() {
        return this.host;
    }
    public String password() {
        return this.password;
    }
    public String port() {
        return this.port;
    }
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConnectionStringConnectionParams defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String database;
        private String host;
        private String password;
        private String port;
        private String username;
        public Builder() {}
        public Builder(GetConnectionStringConnectionParams defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.database = defaults.database;
    	      this.host = defaults.host;
    	      this.password = defaults.password;
    	      this.port = defaults.port;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder database(String database) {
            this.database = Objects.requireNonNull(database);
            return this;
        }
        @CustomType.Setter
        public Builder host(String host) {
            this.host = Objects.requireNonNull(host);
            return this;
        }
        @CustomType.Setter
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
            return this;
        }
        @CustomType.Setter
        public Builder port(String port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public GetConnectionStringConnectionParams build() {
            final var o = new GetConnectionStringConnectionParams();
            o.database = database;
            o.host = host;
            o.password = password;
            o.port = port;
            o.username = username;
            return o;
        }
    }
}
