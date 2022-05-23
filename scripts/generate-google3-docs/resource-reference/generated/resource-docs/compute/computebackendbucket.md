{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}ComputeBackendBucket{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Compute Engine</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/compute/docs/">/compute/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1.backendBucket</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/compute/docs/reference/rest/v1/backendBuckets">/compute/docs/reference/rest/v1/backendBuckets</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpcomputebackendbucket<br>gcpcomputebackendbuckets<br>computebackendbucket</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>compute.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>computebackendbuckets.compute.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/project-id</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
  ```yaml
  bucketRef:
    external: string
    name: string
    namespace: string
  cdnPolicy:
    cacheMode: string
    clientTtl: integer
    defaultTtl: integer
    maxTtl: integer
    negativeCaching: boolean
    negativeCachingPolicy:
    - code: integer
      ttl: integer
    serveWhileStale: integer
    signedUrlCacheMaxAgeSec: integer
  customResponseHeaders:
  - string
  description: string
  edgeSecurityPolicy: string
  enableCdn: boolean
  resourceID: string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>bucketRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Reference to the bucket.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>bucketRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `name` field of a `StorageBucket` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>bucketRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>bucketRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Cloud CDN configuration for this Backend Bucket.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.cacheMode</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Specifies the cache setting for all responses from this backend.
The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"].{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.clientTtl</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Specifies the maximum allowed TTL for cached content served by this origin.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.defaultTtl</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Specifies the default TTL for cached content served by this origin for responses
that do not have an existing valid TTL (max-age or s-max-age).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.maxTtl</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Specifies the maximum allowed TTL for cached content served by this origin.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.negativeCaching</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.negativeCachingPolicy</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.negativeCachingPolicy[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.negativeCachingPolicy[].code</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
can be specified as values, and you cannot specify a status code more than once.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.negativeCachingPolicy[].ttl</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s
(30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.serveWhileStale</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>cdnPolicy.signedUrlCacheMaxAgeSec</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Maximum number of seconds the response to a signed URL request will
be considered fresh. After this time period,
the response will be revalidated before being served.
When serving responses to signed URL requests,
Cloud CDN will internally behave as though
all responses from this backend had a "Cache-Control: public,
max-age=[TTL]" header, regardless of any existing Cache-Control
header. The actual headers served in responses will not be altered.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>customResponseHeaders</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Headers that the HTTP/S load balancer should add to proxied responses.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>customResponseHeaders[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}An optional textual description of the resource; provided by the
client when the resource is created.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>edgeSecurityPolicy</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The security policy associated with this backend bucket.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>enableCdn</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}If true, enable Cloud CDN for this BackendBucket.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
  ```yaml
  conditions:
  - lastTransitionTime: string
    message: string
    reason: string
    status: string
    type: string
  creationTimestamp: string
  observedGeneration: integer
  selfLink: string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>creationTimestamp</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Creation timestamp in RFC3339 text format.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>selfLink</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Basic Backend Bucket
  ```yaml
  # Copyright 2020 Google LLC
  #
  # Licensed under the Apache License, Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at
  #
  #     http://www.apache.org/licenses/LICENSE-2.0
  #
  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeBackendBucket
  metadata:
    name: computebackendbucket-sample-basic
    labels:
      label-one: "value-one"
  spec:
    bucketRef:
      name: ${PROJECT_ID?}-backendbucket-dep-basic
    description: contains a reference to a bucket for use with HTTP(S) load-balancing
  ---
  apiVersion: storage.cnrm.cloud.google.com/v1beta1
  kind: StorageBucket
  metadata:
    # StorageBucket names must be globally unique. Replace ${PROJECT_ID?} with your project ID.
    name: ${PROJECT_ID?}-backendbucket-dep-basic
  ```

### Cdn Enabled Backend Bucket
  ```yaml
  # Copyright 2020 Google LLC
  #
  # Licensed under the Apache License, Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at
  #
  #     http://www.apache.org/licenses/LICENSE-2.0
  #
  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeBackendBucket
  metadata:
    name: computebackendbucket-sample-cdnenabled
    labels:
      label-one: "value-one"
  spec:
    bucketRef:
      name: ${PROJECT_ID?}-backendbucket-dep-cdn
    description: contains a reference to a bucket for use with HTTP(S) load-balancing and integrated CDN, caching on endpoints for only 1/10th the default time
    enableCdn: true
    cdnPolicy:
      signedUrlCacheMaxAgeSec: 360
  ---
  apiVersion: storage.cnrm.cloud.google.com/v1beta1
  kind: StorageBucket
  metadata:
    # StorageBucket names must be globally unique. Replace ${PROJECT_ID?} with your project ID.
    name: ${PROJECT_ID?}-backendbucket-dep-cdn
  ```


{% endblock %}