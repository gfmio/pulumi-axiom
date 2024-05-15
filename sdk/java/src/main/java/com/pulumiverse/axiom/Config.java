// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.gfm.pulumiaxiom;

import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("axiom");
/**
 * The Axiom API token.
 * 
 */
    public String apiToken() {
        return Codegen.stringProp("apiToken").config(config).require();
    }
/**
 * The base url of the axiom api.
 * 
 */
    public Optional<String> baseUrl() {
        return Codegen.stringProp("baseUrl").config(config).get();
    }
}
