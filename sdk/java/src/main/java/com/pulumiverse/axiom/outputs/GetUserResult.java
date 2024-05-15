// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetUserResult {
    /**
     * @return Users email
     * 
     */
    private String email;
    /**
     * @return Users identifier
     * 
     */
    private String id;
    /**
     * @return Users name
     * 
     */
    private String name;
    /**
     * @return Users role
     * 
     */
    private String role;

    private GetUserResult() {}
    /**
     * @return Users email
     * 
     */
    public String email() {
        return this.email;
    }
    /**
     * @return Users identifier
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Users name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Users role
     * 
     */
    public String role() {
        return this.role;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String email;
        private String id;
        private String name;
        private String role;
        public Builder() {}
        public Builder(GetUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.email = defaults.email;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.role = defaults.role;
        }

        @CustomType.Setter
        public Builder email(String email) {
            if (email == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "email");
            }
            this.email = email;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            if (role == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "role");
            }
            this.role = role;
            return this;
        }
        public GetUserResult build() {
            final var _resultValue = new GetUserResult();
            _resultValue.email = email;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.role = role;
            return _resultValue;
        }
    }
}
